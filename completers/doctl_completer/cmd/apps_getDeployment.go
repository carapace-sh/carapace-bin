package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Get a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_getDeploymentCmd).Standalone()
	apps_getDeploymentCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Cause`, `Progress`, `Created`, `Updated`")
	apps_getDeploymentCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_getDeploymentCmd)

	carapace.Gen(apps_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Cause", "Progress", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional
}
