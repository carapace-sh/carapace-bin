package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timingsCmd = &cobra.Command{
	Use:   "timings",
	Short: "Prints timings of all active modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timingsCmd).Standalone()

	timingsCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	timingsCmd.Flags().BoolP("help", "h", false, "Print help")
	timingsCmd.Flags().StringP("jobs", "j", "0", "The number of currently running jobs")
	timingsCmd.Flags().StringP("keymap", "k", "viins", "The keymap of fish/zsh/cmd")
	timingsCmd.Flags().StringP("logical-path", "P", "", "The logical path that the prompt should render for. This path should be a virtual/logical representation of the PATH argument")
	timingsCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	timingsCmd.Flags().StringSlice("pipestatus", nil, "Bash, Fish and Zsh support returning codes for each process in a pipeline")
	timingsCmd.Flags().String("shlvl", "", "The current value of SHLVL, for shells that mis-handle it in $()")
	timingsCmd.Flags().StringP("status", "s", "", "The status code of the previously run command as an unsigned or signed 32bit integer")
	timingsCmd.Flags().StringP("terminal-width", "w", "80", "The width of the current interactive terminal")
	rootCmd.AddCommand(timingsCmd)

	carapace.Gen(timingsCmd).FlagCompletion(carapace.ActionMap{
		"logical-path": carapace.ActionFiles(),
		"path":         carapace.ActionFiles(),
	})
}
