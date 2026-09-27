package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_attachment_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an attachment for a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_attachment_deleteCmd).Standalone()

	volume_attachmentCmd.AddCommand(volume_attachment_deleteCmd)
}
