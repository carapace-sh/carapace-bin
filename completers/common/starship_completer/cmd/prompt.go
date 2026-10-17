package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prints the full starship prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptCmd).Standalone()

	promptCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	promptCmd.Flags().Bool("continuation", false, "Print the continuation prompt (instead of the standard left prompt)")
	promptCmd.Flags().BoolP("help", "h", false, "Print help")
	promptCmd.Flags().StringP("jobs", "j", "0", "The number of currently running jobs")
	promptCmd.Flags().StringP("keymap", "k", "viins", "The keymap of fish/zsh/cmd")
	promptCmd.Flags().StringP("logical-path", "P", "", "The logical path that the prompt should render for. This path should be a virtual/logical representation of the PATH argument")
	promptCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	promptCmd.Flags().StringSlice("pipestatus", nil, "Bash, Fish and Zsh support returning codes for each process in a pipeline")
	promptCmd.Flags().String("profile", "", "Print the prompt with the specified profile name (instead of the standard left prompt)")
	promptCmd.Flags().Bool("right", false, "Print the right prompt (instead of the standard left prompt)")
	promptCmd.Flags().String("shlvl", "", "The current value of SHLVL, for shells that mis-handle it in $()")
	promptCmd.Flags().StringP("status", "s", "", "The status code of the previously run command as an unsigned or signed 32bit integer")
	promptCmd.Flags().StringP("terminal-width", "w", "80", "The width of the current interactive terminal")
	rootCmd.AddCommand(promptCmd)

	carapace.Gen(promptCmd).FlagCompletion(carapace.ActionMap{
		"logical-path": carapace.ActionFiles(),
		"path":         carapace.ActionFiles(),
	})
}
