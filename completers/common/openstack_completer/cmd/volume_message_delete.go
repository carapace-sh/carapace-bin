package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_message_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a volume failure message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_message_deleteCmd).Standalone()

	volume_messageCmd.AddCommand(volume_message_deleteCmd)
}
