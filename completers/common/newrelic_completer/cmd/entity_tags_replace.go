package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_tags_replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace tag:value pairs for the given entity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_tags_replaceCmd).Standalone()
	entity_tags_replaceCmd.Flags().StringP("guid", "g", "", "the entity GUID to replace tag values on")
	entity_tags_replaceCmd.Flags().StringSliceP("tag", "t", nil, "the tag names to replace on the entity")
	entity_tagsCmd.AddCommand(entity_tags_replaceCmd)
}
