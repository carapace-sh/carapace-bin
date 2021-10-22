package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_getCmd).Standalone()
	apps_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Name`, `DefaultIngress`, `ActiveDeployment.ID`, `InProgressDeployment.ID`, `Created`, `Updated`")
	apps_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_getCmd)

	carapace.Gen(apps_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Name", "DefaultIngress", "ActiveDeployment.ID", "InProgressDeployment.ID", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO app-id completion
}
