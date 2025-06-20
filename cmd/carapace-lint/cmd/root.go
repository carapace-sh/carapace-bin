package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}

func Lint(path string) error {
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

	r := regexp.MustCompile(`\.Flags\(\)\.(Bool|String|Float|Int|Uint|Count)[^(]*\("(?P<name>[^"]+)"`)

	previous := ""
	line := 0
	for scanner.Scan() {
		if !r.MatchString(scanner.Text()) {
			previous = ""
		} else {
			matches := r.FindStringSubmatch(scanner.Text())
			current := matches[2]
			if previous != "" && previous > current {
				return fmt.Errorf("%v [%v]: flag '%v' should be before '%v'", path, line, current, previous)
			}
			previous = current
		}
		line += 1
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
