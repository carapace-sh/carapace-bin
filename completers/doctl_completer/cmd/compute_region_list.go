package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_region_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List datacenter regions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_region_listCmd).Standalone()
	compute_region_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Slug`, `Name`, `Available`")
	compute_region_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_regionCmd.AddCommand(compute_region_listCmd)

	carapace.Gen(compute_region_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Slug", "Name", "Available").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
