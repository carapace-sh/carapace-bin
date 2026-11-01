package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wimCmd = &cobra.Command{
	Use:   "wim",
	Short: "WIM-backed file management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wimCmd).Standalone()
	rootCmd.AddCommand(wimCmd)
}
