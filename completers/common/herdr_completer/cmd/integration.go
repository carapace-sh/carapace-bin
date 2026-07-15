package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "Manage built-in agent integrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integrationCmd).Standalone()

	rootCmd.AddCommand(integrationCmd)
}
