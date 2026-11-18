package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Summarise Homebrew's build environment as a plain list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentCmd).Standalone()

	environmentCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	environmentCmd.Flags().BoolP("help", "h", false, "Show this message.")
	environmentCmd.Flags().Bool("plain", false, "Generate plain output even when piped.")
	environmentCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	environmentCmd.Flags().String("shell", "", "Generate a list of environment variables for the specified shell, or `--shell=auto` to detect the current shell.")
	environmentCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(environmentCmd)

	carapace.Gen(environmentCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("auto", "bash", "csh", "fish", "nu", "pwsh", "tcsh", "zsh"),
	})
}
