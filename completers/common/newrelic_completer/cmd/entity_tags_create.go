package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_tags_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create tag:value pairs for the given entity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_tags_createCmd).Standalone()
	entity_tags_createCmd.Flags().StringP("guid", "g", "", "the entity GUID to create tag values on")
	entity_tags_createCmd.Flags().StringSliceP("tag", "t", nil, "the tag names to add to the entity")
	entity_tagsCmd.AddCommand(entity_tags_createCmd)
}
