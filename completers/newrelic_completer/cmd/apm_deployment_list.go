package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apm_deployment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List New Relic APM deployments for an application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_deployment_listCmd).Standalone()
	apm_deploymentCmd.AddCommand(apm_deployment_listCmd)
}
