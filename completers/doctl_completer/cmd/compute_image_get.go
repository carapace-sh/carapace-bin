package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_image_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_getCmd).Standalone()
	compute_image_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_image_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_imageCmd.AddCommand(compute_image_getCmd)

	carapace.Gen(compute_image_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
