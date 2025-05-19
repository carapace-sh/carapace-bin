package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_deployment_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a New Relic entity deployment marker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_deployment_createCmd).Standalone()
	entity_deployment_createCmd.Flags().String("changelog", "", "a URL for the changelog or list of changes if not linkable")
	entity_deployment_createCmd.Flags().String("commit", "", "the commit identifier, for example, a Git commit SHA")
	entity_deployment_createCmd.Flags().String("deepLink", "", "a link back to the system generating the deployment")
	entity_deployment_createCmd.Flags().String("deploymentType", "", "type of deployment, one of BASIC, BLUE_GREEN, CANARY, OTHER, ROLLING or SHADOW")
	entity_deployment_createCmd.Flags().String("description", "", "a description of the deployment")
	entity_deployment_createCmd.Flags().String("groupId", "", "string that can be used to correlate two or more events")
	entity_deployment_createCmd.Flags().StringP("guid", "g", "", "the GUID of the entity associated with this deployment. guid is required.")
	entity_deployment_createCmd.Flags().Int64P("timestamp", "t", 0, "the start time of the deployment, the number of milliseconds since the Unix epoch, defaults to now")
	entity_deployment_createCmd.Flags().StringP("user", "u", "", "username of the deployer or bot")
	entity_deployment_createCmd.Flags().StringP("version", "v", "", "the version of the deployed software, for example, something like v1.1. version is required.")
	entity_deploymentCmd.AddCommand(entity_deployment_createCmd)

	// TODO complete more flags
	carapace.Gen(entity_deployment_createCmd).FlagCompletion(carapace.ActionMap{
		"deploymentType": carapace.ActionValues("BASIC", "BLUE_GREEN", "CANARY", "OTHER", "ROLLING", "SHADOW"),
	})
}
