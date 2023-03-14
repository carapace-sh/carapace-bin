package bridge

import (
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

// ActionCobra bridges https://github.com/spf13/cobra
//
//	var rootCmd = &cobra.Command{
//		Use:                "podman",
//	   Short:              "Simple management tool for pods, containers and images",
//	   Long:               "https://podman.io/",
//		Run:                func(cmd *cobra.Command, args []string) {},
//		DisableFlagParsing: true,
//	}
//
//	func Execute() error {
//		return rootCmd.Execute()
//	}
//	func init() {
//		carapace.Gen(rootCmd).Standalone()
//
//		carapace.Gen(rootCmd).PositionalAnyCompletion(
//			cobracomplete.ActionCobra("podman"),
//		)
//	}
func ActionCobra(command ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(command) == 0 {
			return carapace.ActionMessage("missing argument [ActionCobra]")
		}

		args := []string{"__complete"}
		args = append(args, command[1:]...)
		args = append(args, c.Args...)
		args = append(args, c.CallbackValue)
		return carapace.ActionExecCommand(command[0], args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			if len(lines) == 0 {
				return carapace.ActionMessage("unexpected command output")
			}

			// TODO experimental - directives not yet fully tested
			var action carapace.Action
			directive, err := readDirective(lines)
			if err != nil {
				return carapace.ActionValues()
			}

			if directive.matches(cobra.ShellCompDirectiveFilterDirs) {
				return actionDirectories(lines)
			} else if directive.matches(cobra.ShellCompDirectiveFilterFileExt) {
				return actionFiles(lines)
			} else {
				action = actionValues(lines)
			}

			if len(lines) < 3 && !directive.matches(cobra.ShellCompDirectiveNoFileComp) {
				action = carapace.ActionFiles()
			}

			if directive.matches(cobra.ShellCompDirectiveError) {
				action = carapace.ActionValues()
			}

			if directive.matches(cobra.ShellCompDirectiveNoSpace) {
				action = action.NoSpace()
			}
			return action
		})
	})
}

type compDirective int

func (d compDirective) matches(cobraDirective cobra.ShellCompDirective) bool {
	return d&compDirective(cobraDirective) != 0
}

func readDirective(lines []string) (compDirective, error) {
	value, err := strconv.Atoi(strings.TrimPrefix(lines[len(lines)-2], ":"))
	if err != nil {
		return -1, err
	}
	return compDirective(value), nil
}

func actionValues(lines []string) carapace.Action {
	vals := make([]string, 0)
	for _, line := range lines[:len(lines)-2] {
		if splitted := strings.SplitN(line, "\t", 2); len(splitted) == 2 {
			vals = append(vals, splitted...)
		} else if len(splitted) == 1 {
			vals = append(vals, splitted[0], "")
		}
	}
	return carapace.ActionValuesDescribed(vals...)
}

func actionFiles(lines []string) carapace.Action {
	extensions := make([]string, 0)
	for _, line := range lines[:len(lines)-2] {
		extensions = append(extensions, "."+line)
	}
	return carapace.ActionFiles(extensions...)
}

func actionDirectories(lines []string) carapace.Action {
	if len(lines) > 2 {
		return carapace.ActionDirectories().Chdir(lines[0])
	} else {
		return carapace.ActionDirectories()
	}
}
