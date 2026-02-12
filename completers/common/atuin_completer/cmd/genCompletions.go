package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genCompletionsCmd = &cobra.Command{
	Use:   "gen-completions",
	Short: "Generate shell completions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genCompletionsCmd).Standalone()

	genCompletionsCmd.Flags().BoolP("help", "h", false, "Print help")
	genCompletionsCmd.Flags().StringP("out-dir", "o", "", "Set the output directory")
	genCompletionsCmd.Flags().StringP("shell", "s", "", "Set the shell for generating completions")
	genCompletionsCmd.MarkFlagRequired("shell")
	rootCmd.AddCommand(genCompletionsCmd)

	carapace.Gen(genCompletionsCmd).FlagCompletion(carapace.ActionMap{
		"out-dir": carapace.ActionDirectories(),
		"shell":   carapace.ActionValues("bash", "elvish", "fish", "nushell", "powershell", "zsh"),
	})
}
