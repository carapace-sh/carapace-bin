package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backendCmd).Standalone()

	volumeCmd.AddCommand(volume_backendCmd)
}
