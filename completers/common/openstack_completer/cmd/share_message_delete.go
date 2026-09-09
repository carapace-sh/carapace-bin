package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_message_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove one or more messages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_message_deleteCmd).Standalone()

	share_messageCmd.AddCommand(share_message_deleteCmd)
}
