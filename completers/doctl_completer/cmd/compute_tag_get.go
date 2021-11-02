package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_tag_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_tag_getCmd).Standalone()
	compute_tag_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `DropletCount`")
	compute_tag_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_tagCmd.AddCommand(compute_tag_getCmd)

	carapace.Gen(compute_tag_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Name", "DropletCount").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
