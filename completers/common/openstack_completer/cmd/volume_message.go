package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_messageCmd = &cobra.Command{
	Use:   "message",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_messageCmd).Standalone()

	volumeCmd.AddCommand(volume_messageCmd)
}
