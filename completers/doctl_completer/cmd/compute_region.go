package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_regionCmd = &cobra.Command{
	Use:   "region",
	Short: "Display commands to list datacenter regions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_regionCmd).Standalone()
	computeCmd.AddCommand(compute_regionCmd)
}
