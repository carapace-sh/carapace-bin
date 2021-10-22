package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_kernelsCmd = &cobra.Command{
	Use:   "kernels",
	Short: "List available Droplet kernels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_kernelsCmd).Standalone()
	compute_droplet_kernelsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Version`")
	compute_droplet_kernelsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletCmd.AddCommand(compute_droplet_kernelsCmd)

	carapace.Gen(compute_droplet_kernelsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Version").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
