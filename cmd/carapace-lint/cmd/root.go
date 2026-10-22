package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/owenrumney/go-sarif/sarif"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace-lint",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := 0
		var issues []lintIssue
		for _, arg := range args {
			lintIssues, err := Lint(arg)
			if err != nil {
				println(err.Error())
				exitCode = 1
				continue
			}
			issues = append(issues, lintIssues...)

			flagActionIssues, err := LintFlagActions(arg)
			if err != nil {
				println(err.Error())
				exitCode = 1
				continue
			}
			issues = append(issues, flagActionIssues...)
		}

		if len(issues) > 0 {
			exitCode = 1
		}

		if sarifPath != "" {
			if err := writeSarif(sarifPath, issues); err != nil {
				println(err.Error())
				os.Exit(1)
			}
		} else {
			for _, issue := range issues {
				println(fmt.Sprintf("%s [%d]: %s", issue.file, issue.line, issue.message))
			}
		}
		os.Exit(exitCode)
	},
}

var fixFlagsOrder bool
var sarifPath string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVar(&fixFlagsOrder, "fix-flags-order", false, "auto-fix flags order")
	rootCmd.Flags().StringVar(&sarifPath, "sarif", "", "output sarif report to file")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}

type (
	sourceLine struct {
		Source     string
		LineNumber int
		FlagName   string
	}

	flagsBlockDef struct {
		Start int // inclusive
		End   int // exclusive
	}
)

func Lint(path string) ([]lintIssue, error) {
	if !strings.HasSuffix(path, ".go") {
		return nil, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`\.Flags\(\)\.(Bool|String|Float|Int|Uint|Count)[^(]*\("(?P<name>[^"]+)"`)

	contentLines := bytes.Split(content, []byte{'\n'})
	lines := make([]sourceLine, len(contentLines))

	for i, line := range contentLines {
		lineSource := string(line)

		if matches := r.FindStringSubmatch(lineSource); matches != nil {
			lines[i] = sourceLine{
				LineNumber: i + 1,
				Source:     lineSource,
				FlagName:   matches[2],
			}
		} else {
			lines[i] = sourceLine{
				LineNumber: i + 1,
				Source:     lineSource,
			}
		}
	}

	var defs []flagsBlockDef
	for i := 0; i < len(lines); {
		// regular line, do nothing
		if !lines[i].isFlagLine() {
			i++
			continue
		}

		// `i` is the start of a contiguous "flags block".
		// we now have to find it's end
		j := i + 1
		for j < len(lines) && lines[j].isFlagLine() {
			j++
		}

		// the flags block consists of only one flag line.
		// no need to sort it
		if j-i == 1 {
			i++
			continue
		}

		defs = append(defs, flagsBlockDef{
			Start: i,
			End:   j,
		})

		// we know that `j` is no flag line and can safely skip it
		i = j + 1
	}

	if fixFlagsOrder {
		for _, def := range defs {
			sort.Slice(lines[def.Start:def.End], func(a, b int) bool {
				return lines[def.Start+a].FlagName < lines[def.Start+b].FlagName
			})
		}

		var buf bytes.Buffer
		for i, line := range lines {
			buf.WriteString(line.Source)
			isLastLine := i == len(lines)-1
			if !isLastLine {
				buf.WriteByte('\n')
			}
		}

		return nil, os.WriteFile(path, buf.Bytes(), 0644)
	} else {
		var issues []lintIssue
		for _, def := range defs {
			block := lines[def.Start:def.End]
			for i := 1; i < len(block); i++ {
				prev := block[i-1]
				current := block[i]

				if current.FlagName < prev.FlagName {
					issues = append(issues, lintIssue{
						file:    path,
						line:    current.LineNumber,
						message: fmt.Sprintf("flag '%s' should be before '%s'", current.FlagName, prev.FlagName),
						ruleID:  "carapace-lint/flags-order",
					})
				}
			}
		}
		return issues, nil
	}
}

func LintFlagActions(path string) ([]lintIssue, error) {
	if !strings.HasSuffix(path, ".go") {
		return nil, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rStart := regexp.MustCompile(`^\tcarapace\.Gen\([^)]+\)\.FlagCompletion\(carapace.ActionMap{$`)
	rFlagName := regexp.MustCompile(`^\t\t"(?P<name>[^"]+)":.*$`)
	rEnd := regexp.MustCompile(`^\t}\)?$`)

	var issues []lintIssue
	line := 0
	previous := ""
outer:
	for scanner.Scan() {
		if rStart.MatchString(scanner.Text()) {
			line += 2
			for scanner.Scan() {
				if rEnd.MatchString(scanner.Text()) {
					break outer
				}
				if rFlagName.MatchString(scanner.Text()) {
					matches := rFlagName.FindStringSubmatch(scanner.Text())
					current := matches[1]
					if previous != "" && previous > current {
						issues = append(issues, lintIssue{
							file:    path,
							line:    line,
							message: fmt.Sprintf("flag completion '%s' should be before '%s'", current, previous),
							ruleID:  "carapace-lint/flag-actions-order",
						})
					}
					previous = current
				}
				line += 1
			}
		}
		line += 1
	}
	return issues, nil
}

func (l sourceLine) isFlagLine() bool {
	return l.FlagName != ""
}

type lintIssue struct {
	file    string
	line    int
	message string
	ruleID  string
}

func writeSarif(path string, issues []lintIssue) error {
	report, err := sarif.New(sarif.Version210)
	if err != nil {
		return err
	}

	run := sarif.NewRun("carapace-lint", "https://github.com/carapace-sh/carapace-bin")
	run.AddRule("carapace-lint/flags-order").WithDescription("Flags should be in alphabetical order")
	run.AddRule("carapace-lint/flag-actions-order").WithDescription("Flag completions should be in alphabetical order")

	for _, issue := range issues {
		result := run.AddResult(issue.ruleID)
		result.WithLevel("warning").
			WithMessage(sarif.NewTextMessage(issue.message))

		if issue.file != "" && issue.line > 0 {
			result.WithLocation(sarif.NewLocationWithPhysicalLocation(
				sarif.NewPhysicalLocation().
					WithArtifactLocation(sarif.NewSimpleArtifactLocation(issue.file)).
					WithRegion(sarif.NewSimpleRegion(issue.line, issue.line)),
			))
		}
	}

	report.AddRun(run)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return report.PrettyWrite(file)
}
