package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "List available Droplet sizes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sizeCmd).Standalone()
	computeCmd.AddCommand(compute_sizeCmd)
}
