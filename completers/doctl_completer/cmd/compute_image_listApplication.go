package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_image_listApplicationCmd = &cobra.Command{
	Use:   "list-application",
	Short: "List available One-Click Apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_listApplicationCmd).Standalone()
	compute_image_listApplicationCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_image_listApplicationCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_image_listApplicationCmd.Flags().Bool("public", true, "List public images")
	compute_imageCmd.AddCommand(compute_image_listApplicationCmd)

	carapace.Gen(compute_image_listApplicationCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
