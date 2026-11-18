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

	shCmd.Flags().String("cmd", "", "Execute commands in a non-interactive shell.")
	shCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	shCmd.Flags().String("env", "", "Use the standard `PATH` instead of superenv's when `std` is passed.")
	shCmd.Flags().BoolP("help", "h", false, "Show this message.")
	shCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	shCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(shCmd)

	carapace.Gen(shCmd).FlagCompletion(carapace.ActionMap{
		"env": carapace.ActionValues("std"),
	})
}
