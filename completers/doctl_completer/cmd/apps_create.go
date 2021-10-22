package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_createCmd).Standalone()
	apps_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Name`, `DefaultIngress`, `ActiveDeployment.ID`, `InProgressDeployment.ID`, `Created`, `Updated`")
	apps_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	apps_createCmd.Flags().String("spec", "", "Path to an app spec in JSON or YAML format. Set to \"-\" to read from stdin. (required)")
	apps_createCmd.Flags().Bool("upsert", false, "Boolean that specifies whether the app should be updated if it already exists")
	apps_createCmd.Flags().Bool("wait", false, "Boolean that specifies whether to wait for an app to complete before returning control to the terminal")
	appsCmd.AddCommand(apps_createCmd)

	carapace.Gen(apps_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Name", "DefaultIngress", "ActiveDeployment.ID", "InProgressDeployment.ID", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
		"spec": carapace.ActionFiles(),
	})

	// TODO app id postional completion
}
