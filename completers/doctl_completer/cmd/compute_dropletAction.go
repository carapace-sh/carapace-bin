package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletActionCmd = &cobra.Command{
	Use:   "droplet-action",
	Short: "Display Droplet action commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletActionCmd).Standalone()
	computeCmd.AddCommand(compute_dropletActionCmd)
}
