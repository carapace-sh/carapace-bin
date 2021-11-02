package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_image_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an image's metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_updateCmd).Standalone()
	compute_image_updateCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_image_updateCmd.Flags().String("image-name", "", "Image name (required)")
	compute_image_updateCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_imageCmd.AddCommand(compute_image_updateCmd)

	carapace.Gen(compute_image_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
