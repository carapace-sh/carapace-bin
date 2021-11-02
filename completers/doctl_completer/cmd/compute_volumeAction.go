package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volumeActionCmd = &cobra.Command{
	Use:   "volume-action",
	Short: "Display commands to perform actions on a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeActionCmd).Standalone()
	computeCmd.AddCommand(compute_volumeActionCmd)
}
