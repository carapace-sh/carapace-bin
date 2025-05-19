package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shCmd = &cobra.Command{
	Use:     "sh",
	Short:   "Enter an interactive shell for Homebrew's build environment",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shCmd).Standalone()

	shCmd.Flags().Bool("cmd", false, "Execute commands in a non-interactive shell.")
	shCmd.Flags().Bool("debug", false, "Display any debugging information.")
	shCmd.Flags().Bool("env", false, "Use the standard `PATH` instead of superenv's when `std` is passed.")
	shCmd.Flags().Bool("help", false, "Show this message.")
	shCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	shCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(shCmd)
}
