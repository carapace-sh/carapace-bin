package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rasCmd = &cobra.Command{
	Use:   "ras",
	Short: "Remote Access configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rasCmd).Standalone()
	rootCmd.AddCommand(rasCmd)
}
