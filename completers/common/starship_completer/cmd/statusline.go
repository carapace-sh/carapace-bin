package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statuslineCmd = &cobra.Command{
	Use:   "statusline",
	Short: "Prints the statusline with a specific profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statuslineCmd).Standalone()

	statuslineCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	statuslineCmd.Flags().BoolP("help", "h", false, "Print help")
	statuslineCmd.Flags().StringP("jobs", "j", "0", "The number of currently running jobs")
	statuslineCmd.Flags().StringP("keymap", "k", "viins", "The keymap of fish/zsh/cmd")
	statuslineCmd.Flags().StringP("logical-path", "P", "", "The logical path that the prompt should render for. This path should be a virtual/logical representation of the PATH argument")
	statuslineCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	statuslineCmd.Flags().StringSlice("pipestatus", nil, "Bash, Fish and Zsh support returning codes for each process in a pipeline")
	statuslineCmd.Flags().String("profile", "", "")
	statuslineCmd.Flags().String("shlvl", "", "The current value of SHLVL, for shells that mis-handle it in $()")
	statuslineCmd.Flags().StringP("status", "s", "", "The status code of the previously run command as an unsigned or signed 32bit integer")
	statuslineCmd.Flags().StringP("terminal-width", "w", "80", "The width of the current interactive terminal")
	rootCmd.AddCommand(statuslineCmd)

	carapace.Gen(statuslineCmd).FlagCompletion(carapace.ActionMap{
		"logical-path": carapace.ActionFiles(),
		"path":         carapace.ActionFiles(),
	})

	carapace.Gen(statuslineCmd).PositionalCompletion(
		carapace.ActionValues("claude-code"),
	)
}
