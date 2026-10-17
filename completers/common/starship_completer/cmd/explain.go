package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explains the currently showing modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	explainCmd.Flags().BoolP("help", "h", false, "Print help")
	explainCmd.Flags().StringP("jobs", "j", "0", "The number of currently running jobs")
	explainCmd.Flags().StringP("keymap", "k", "viins", "The keymap of fish/zsh/cmd")
	explainCmd.Flags().StringP("logical-path", "P", "", "The logical path that the prompt should render for. This path should be a virtual/logical representation of the PATH argument")
	explainCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	explainCmd.Flags().StringSlice("pipestatus", nil, "Bash, Fish and Zsh support returning codes for each process in a pipeline")
	explainCmd.Flags().String("shlvl", "", "The current value of SHLVL, for shells that mis-handle it in $()")
	explainCmd.Flags().StringP("status", "s", "", "The status code of the previously run command as an unsigned or signed 32bit integer")
	explainCmd.Flags().StringP("terminal-width", "w", "80", "The width of the current interactive terminal")
	rootCmd.AddCommand(explainCmd)

	carapace.Gen(explainCmd).FlagCompletion(carapace.ActionMap{
		"logical-path": carapace.ActionFiles(),
		"path":         carapace.ActionFiles(),
	})
}
