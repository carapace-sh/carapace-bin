package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_listCmd).Standalone()
	apps_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Name`, `DefaultIngress`, `ActiveDeployment.ID`, `InProgressDeployment.ID`, `Created`, `Updated`")
	apps_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_listCmd)

	carapace.Gen(apps_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Name", "DefaultIngress", "ActiveDeployment.ID", "InProgressDeployment.ID", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
