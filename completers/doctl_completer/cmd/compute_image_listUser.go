package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_image_listUserCmd = &cobra.Command{
	Use:   "list-user",
	Short: "List user-created images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_listUserCmd).Standalone()
	compute_image_listUserCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_image_listUserCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_image_listUserCmd.Flags().Bool("public", false, "List public images")
	compute_imageCmd.AddCommand(compute_image_listUserCmd)

	carapace.Gen(compute_image_listUserCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
