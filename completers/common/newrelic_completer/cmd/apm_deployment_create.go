package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apm_deployment_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a New Relic APM deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_deployment_createCmd).Standalone()
	apm_deployment_createCmd.Flags().String("change-log", "", "the change log stored with the deployment")
	apm_deployment_createCmd.Flags().String("description", "", "the description stored with the deployment")
	apm_deployment_createCmd.Flags().StringP("revision", "r", "", "a freeform string representing the revision of the deployment")
	apm_deployment_createCmd.Flags().String("user", "", "the user creating with the deployment")
	apm_deploymentCmd.AddCommand(apm_deployment_createCmd)
}
