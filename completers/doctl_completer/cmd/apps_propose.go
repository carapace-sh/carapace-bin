package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_proposeCmd = &cobra.Command{
	Use:   "propose",
	Short: "Propose an app spec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_proposeCmd).Standalone()
	apps_proposeCmd.Flags().String("app", "", "An optional existing app ID. If specified, the app spec will be treated as a proposed update to the existing app.")
	apps_proposeCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Name`, `DefaultIngress`, `ActiveDeployment.ID`, `InProgressDeployment.ID`, `Created`, `Updated`")
	apps_proposeCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	apps_proposeCmd.Flags().String("spec", "", "Path to an app spec in JSON or YAML format.")
	appsCmd.AddCommand(apps_proposeCmd)

	// TODO app
	carapace.Gen(apps_proposeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Name", "DefaultIngress", "ActiveDeployment.ID", "InProgressDeployment.ID", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
		"spec": carapace.ActionFiles(),
	})

	// TODO positional
}
