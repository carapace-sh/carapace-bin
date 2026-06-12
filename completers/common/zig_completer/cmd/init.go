package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [options]",
	Short:   "Initialize a Zig package in the current directory",
	GroupID: "project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("help", "h", false, "Print help")
	initCmd.Flags().BoolP("minimal", "m", false, "Use minimal init template")
}
