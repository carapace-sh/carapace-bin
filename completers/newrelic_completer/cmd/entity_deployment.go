package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Manage deployment markers for a New Relic entity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_deploymentCmd).Standalone()
	entityCmd.AddCommand(entity_deploymentCmd)
}
