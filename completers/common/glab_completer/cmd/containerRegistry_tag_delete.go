package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerRegistry_tag_deleteCmd = &cobra.Command{
	Use:     "delete <repository-id> [<tag-name>] [flags]",
	Short:   "Delete container registry tags.",
	Aliases: []string{"del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_tag_deleteCmd).Standalone()

	containerRegistry_tag_deleteCmd.Flags().String("keep-n", "", "Keep the latest N matching tags. Bulk deletion only; scheduled asynchronously.")
	containerRegistry_tag_deleteCmd.Flags().String("name-regex-delete", "", "Regular expression for tag names to delete. Bulk deletion only; scheduled asynchronously.")
	containerRegistry_tag_deleteCmd.Flags().String("name-regex-keep", "", "Regular expression for tag names to keep. Bulk deletion only; scheduled asynchronously.")
	containerRegistry_tag_deleteCmd.Flags().String("older-than", "", "Delete tags older than the given duration, such as 7d or 1month. Bulk deletion only; scheduled asynchronously.")
	containerRegistry_tag_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	containerRegistry_tagCmd.AddCommand(containerRegistry_tag_deleteCmd)
}
