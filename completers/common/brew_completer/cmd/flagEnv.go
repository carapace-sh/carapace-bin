package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Summarise Homebrew's build environment as a plain list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagEnvCmd).Standalone()

	flagEnvCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagEnvCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagEnvCmd.Flags().Bool("plain", false, "Generate plain output even when piped.")
	flagEnvCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagEnvCmd.Flags().String("shell", "", "Generate a list of environment variables for the specified shell, or `--shell=auto` to detect the current shell.")
	flagEnvCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")

	carapace.Gen(flagEnvCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("auto", "bash", "csh", "fish", "nu", "pwsh", "tcsh", "zsh"),
	})
}
