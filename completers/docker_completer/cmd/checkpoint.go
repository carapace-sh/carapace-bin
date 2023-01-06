package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkpointCmd = &cobra.Command{
	Use:     "checkpoint",
	Short:   "Manage checkpoints",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkpointCmd).Standalone()
	rootCmd.AddCommand(checkpointCmd)
}
