package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_image_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List images on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_listCmd).Standalone()
	compute_image_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_image_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_image_listCmd.Flags().Bool("public", false, "List public images")
	compute_imageCmd.AddCommand(compute_image_listCmd)

	carapace.Gen(compute_image_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
