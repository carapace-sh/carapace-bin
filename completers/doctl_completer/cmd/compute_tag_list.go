package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_tag_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_tag_listCmd).Standalone()
	compute_tag_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `DropletCount`")
	compute_tag_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_tagCmd.AddCommand(compute_tag_listCmd)

	carapace.Gen(compute_tag_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Name", "DropletCount").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
