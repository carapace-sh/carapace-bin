package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Zig package in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("strip", false, "Omit debug symbols")
	initCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(initCmd)
}
