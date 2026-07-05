package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hibernateCmd = &cobra.Command{
	Use:   "hibernate",
	Short: "enable or disable hibernate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hibernateCmd).Standalone()
	rootCmd.AddCommand(hibernateCmd)
}
