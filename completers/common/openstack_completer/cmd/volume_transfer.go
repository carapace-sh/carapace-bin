package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_transferCmd).Standalone()

	volumeCmd.AddCommand(volume_transferCmd)
}
