package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_size_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Droplet sizes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_size_listCmd).Standalone()
	compute_size_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Slug`, `Memory`, `VCPUs`, `Disk`, `PriceMonthly`, `PriceHourly`")
	compute_size_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_sizeCmd.AddCommand(compute_size_listCmd)

	carapace.Gen(compute_size_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Slug", "Memory", "VCPUs", "Disk", "PriceMonthly", "PriceHourly").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
