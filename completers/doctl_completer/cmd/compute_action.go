package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Display commands for retrieving resource action history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_actionCmd).Standalone()
	computeCmd.AddCommand(compute_actionCmd)
}
