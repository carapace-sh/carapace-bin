package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Create a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_createDeploymentCmd).Standalone()
	apps_createDeploymentCmd.Flags().Bool("force-rebuild", false, "Force a re-build even if a previous build is eligible for reuse")
	apps_createDeploymentCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Cause`, `Progress`, `Created`, `Updated`")
	apps_createDeploymentCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	apps_createDeploymentCmd.Flags().Bool("wait", false, "Boolean that specifies whether to wait for apps deployment to complete before returning control to the terminal")
	appsCmd.AddCommand(apps_createDeploymentCmd)

	carapace.Gen(apps_createDeploymentCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Cause", "Progress", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
