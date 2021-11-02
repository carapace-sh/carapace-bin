package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_listRegionsCmd = &cobra.Command{
	Use:   "list-regions",
	Short: "List App Platform regions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_listRegionsCmd).Standalone()
	apps_listRegionsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Slug`, `Label`, `Continent`, `DataCenters`, `Disabled`, `Reason`, `Default`")
	apps_listRegionsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_listRegionsCmd)

	carapace.Gen(apps_listRegionsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Slug", "Label", "Continent", "DataCenters", "Disabled", "Reason", "Default").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional
}
