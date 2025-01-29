package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload all volumes from volume plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_reloadCmd).Standalone()

	volumeCmd.AddCommand(volume_reloadCmd)
}
