package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "List all deployments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_listDeploymentsCmd).Standalone()
	apps_listDeploymentsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Cause`, `Progress`, `Created`, `Updated`")
	apps_listDeploymentsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_listDeploymentsCmd)

	carapace.Gen(apps_listDeploymentsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Cause", "Progress", "Created", "Updated").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional
}
