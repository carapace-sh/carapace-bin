package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert a volume to a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_revertCmd).Standalone()

	volumeCmd.AddCommand(volume_revertCmd)
}
