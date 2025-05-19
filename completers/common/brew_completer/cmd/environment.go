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

	environmentCmd.Flags().Bool("debug", false, "Display any debugging information.")
	environmentCmd.Flags().Bool("help", false, "Show this message.")
	environmentCmd.Flags().Bool("plain", false, "Generate plain output even when piped.")
	environmentCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	environmentCmd.Flags().Bool("shell", false, "Generate a list of environment variables for the specified shell, or `--shell=auto` to detect the current shell.")
	environmentCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(environmentCmd)
}
