package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Display commands to manage block storage volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeCmd).Standalone()
	computeCmd.AddCommand(compute_volumeCmd)
}
