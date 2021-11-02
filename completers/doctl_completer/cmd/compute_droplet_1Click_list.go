package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_1Click_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of Droplet 1-Click applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_1Click_listCmd).Standalone()
	compute_droplet_1Click_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `SLUG`, `TYPE`")
	compute_droplet_1Click_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_droplet_1ClickCmd.AddCommand(compute_droplet_1Click_listCmd)

	carapace.Gen(compute_droplet_1Click_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("SLUG", "TYPE").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
