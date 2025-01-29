package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_unmountCmd = &cobra.Command{
	Use:   "unmount NAME",
	Short: "Unmount volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_unmountCmd).Standalone()

	volumeCmd.AddCommand(volume_unmountCmd)
}
