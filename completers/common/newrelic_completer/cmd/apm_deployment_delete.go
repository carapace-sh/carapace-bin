package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apm_deployment_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a New Relic APM deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_deployment_deleteCmd).Standalone()
	apm_deployment_deleteCmd.Flags().IntP("deploymentID", "d", 0, "the ID of the deployment to be deleted")
	apm_deploymentCmd.AddCommand(apm_deployment_deleteCmd)
}
