package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_tags_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the given tag:value pairs from the given entity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_tags_deleteCmd).Standalone()
	entity_tags_deleteCmd.Flags().StringP("guid", "g", "", "the entity GUID to delete tags on")
	entity_tags_deleteCmd.Flags().StringSliceP("tag", "t", nil, "the tag keys to delete from the entity")
	entity_tagsCmd.AddCommand(entity_tags_deleteCmd)
}
