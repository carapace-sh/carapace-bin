package cmd

import (
	"bufio"
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

	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r := regexp.MustCompile(`\.Flags\(\)\.(Bool|String|Float|Int|Uint|Count)[^(]*\("(?P<name>[^"]+)"`)

	lines := []sourceLine{}
	for scanner.Scan() {
		line := scanner.Text()

		if r.MatchString(line) {
			matches := r.FindStringSubmatch(line)

			lines = append(lines, sourceLine{
				LineNumber: len(lines) + 1,
				Source:     line,
				FlagName:   matches[2],
			})
		} else {
			lines = append(lines, sourceLine{
				LineNumber: len(lines) + 1,
				Source:     line,
			})
		}
	}

	defs := []flagsBlockDef{}

	i := 0
	for i < len(lines) {
		// regular line, do nothing
		if !lines[i].isFlagLine() {
			i++
			continue
		}

		// `i` is the start of a contiguous "flags block".
		// we now have to find it's end
		j := i
		for j < len(lines) && lines[j].isFlagLine() {
			j++
		}

		// the flags block consists of only one flag line.
		// there is no need to sort
		if i == j {
			continue
		}

		defs = append(defs, flagsBlockDef{
			Start: i,
			End:   j,
		})
		i = j
	}

	if fixFlagsOrder {
		for _, def := range defs {
			sort.Slice(lines[def.Start:def.End], func(a, b int) bool {
				return lines[def.Start+a].FlagName < lines[def.Start+b].FlagName
			})
		}

		if _, err := file.Seek(0, 0); err != nil {
			return fmt.Errorf("resetting file offset failed: %v", err)
		}

		for i, line := range lines {
			_, _ = file.WriteString(line.Source)

			isLastLine := i == len(lines)-1
			if !isLastLine {
				_, _ = file.WriteString("\n")
			}
		}
	} else {
		for _, def := range defs {
			block := lines[def.Start:def.End]
			for blockIndex := range block {
				if blockIndex == 0 {
					continue
				}

				prev := block[blockIndex-1]
				current := block[blockIndex]

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
