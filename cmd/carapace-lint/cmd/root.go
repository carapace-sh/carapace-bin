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
		for _, arg := range args {
			if err := Lint(arg); err != nil {
				println(err.Error())
				exitCode = 1
			} else if err := LintFlagActions(arg); err != nil {
				println(err.Error())
				exitCode = 1
			}
		}
		os.Exit(exitCode)
	},
}

var fixFlagsOrder bool

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVar(&fixFlagsOrder, "fix-flags-order", false, "auto-fix flags order")

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

func Lint(path string) error {
	if !strings.HasSuffix(path, ".go") {
		return nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
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

		// the flags block consists of only one flag line
		if i == j {
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

		return os.WriteFile(path, buf.Bytes(), 0644)
	} else {
		for _, def := range defs {
			block := lines[def.Start:def.End]
			for i := 1; i < len(block); i++ {
				prev := block[i-1]
				current := block[i]

				if current.FlagName < prev.FlagName {
					return fmt.Errorf("%s [%d]: flag '%s' should be before '%s'", path, current.LineNumber, current.FlagName, prev.FlagName)
				}
			}
		}
	}

	return nil
}

func LintFlagActions(path string) error {
	if !strings.HasSuffix(path, ".go") {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rStart := regexp.MustCompile(`^\tcarapace\.Gen\([^)]+\)\.FlagCompletion\(carapace.ActionMap{$`)
	rFlagName := regexp.MustCompile(`^\t\t"(?P<name>[^"]+)":.*$`)
	rEnd := regexp.MustCompile(`^\t}\)?$`)

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
						return fmt.Errorf("%v [%v]: flag completion '%v' should be before '%v'", path, line, current, previous)
					}
					previous = current
				}
				line += 1
			}
		}
		line += 1
	}
	return nil
}

func (l sourceLine) isFlagLine() bool {
	return l.FlagName != ""
}
