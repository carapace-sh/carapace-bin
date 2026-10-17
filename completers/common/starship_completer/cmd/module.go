package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Prints a specific prompt module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleCmd).Standalone()

	moduleCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	moduleCmd.Flags().BoolP("help", "h", false, "Print help")
	moduleCmd.Flags().StringP("jobs", "j", "0", "The number of currently running jobs")
	moduleCmd.Flags().StringP("keymap", "k", "viins", "The keymap of fish/zsh/cmd")
	moduleCmd.Flags().BoolP("list", "l", false, "List out all supported modules")
	moduleCmd.Flags().StringP("logical-path", "P", "", "The logical path that the prompt should render for. This path should be a virtual/logical representation of the PATH argument")
	moduleCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	moduleCmd.Flags().StringSlice("pipestatus", nil, "Bash, Fish and Zsh support returning codes for each process in a pipeline")
	moduleCmd.Flags().String("shlvl", "", "The current value of SHLVL, for shells that mis-handle it in $()")
	moduleCmd.Flags().StringP("status", "s", "", "The status code of the previously run command as an unsigned or signed 32bit integer")
	moduleCmd.Flags().StringP("terminal-width", "w", "80", "The width of the current interactive terminal")
	rootCmd.AddCommand(moduleCmd)

	carapace.Gen(moduleCmd).FlagCompletion(carapace.ActionMap{
		"logical-path": carapace.ActionFiles(),
		"path":         carapace.ActionFiles(),
	})

	carapace.Gen(moduleCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("starship", "module", "--list")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[2 : len(lines)-1]...)
			})
		}),
	)
}
