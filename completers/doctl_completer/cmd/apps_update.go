package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_updateCmd).Standalone()
	apps_updateCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Name`, `DefaultIngress`, `ActiveDeployment.ID`, `InProgressDeployment.ID`, `Created`, `Updated`")
	apps_updateCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	apps_updateCmd.Flags().String("spec", "", "Path to an app spec in JSON or YAML format. Set to \"-\" to read from stdin. (required)")
	apps_updateCmd.Flags().Bool("wait", false, "Boolean that specifies whether to wait for an app to complete before returning control to the terminal")
	appsCmd.AddCommand(apps_updateCmd)

	carapace.Gen(apps_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Name", "DefaultIngress", "ActiveDeployment.ID", "InProgressDeployment.ID", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
		"spec": carapace.ActionFiles(),
	})

	// TODO positional
}
