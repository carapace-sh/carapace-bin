package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletCmd = &cobra.Command{
	Use:   "droplet",
	Short: "Manage virtual machines (Droplets)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletCmd).Standalone()
	computeCmd.AddCommand(compute_dropletCmd)
}
