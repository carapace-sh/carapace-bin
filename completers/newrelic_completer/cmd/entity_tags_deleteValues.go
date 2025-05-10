package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_tags_deleteValuesCmd = &cobra.Command{
	Use:   "delete-values",
	Short: "Delete the given tag/value pairs from the given entity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_tags_deleteValuesCmd).Standalone()
	entity_tags_deleteValuesCmd.Flags().StringP("guid", "g", "", "the entity GUID to delete tag values on")
	entity_tags_deleteValuesCmd.Flags().StringSliceP("value", "v", nil, "the tag key:value pairs to delete from the entity")
	entity_tagsCmd.AddCommand(entity_tags_deleteValuesCmd)
}
