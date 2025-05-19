package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:     "context",
	Short:   "Manage contexts",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contextCmd).Standalone()

	rootCmd.AddCommand(contextCmd)
}
