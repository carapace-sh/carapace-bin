package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_mountCmd = &cobra.Command{
	Use:   "mount NAME",
	Short: "Mount volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_mountCmd).Standalone()

	volumeCmd.AddCommand(volume_mountCmd)
}
