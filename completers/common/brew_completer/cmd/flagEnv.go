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

	flagEnvCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagEnvCmd.Flags().Bool("help", false, "Show this message.")
	flagEnvCmd.Flags().Bool("plain", false, "Generate plain output even when piped.")
	flagEnvCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagEnvCmd.Flags().Bool("shell", false, "Generate a list of environment variables for the specified shell, or `--shell=auto` to detect the current shell.")
	flagEnvCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
