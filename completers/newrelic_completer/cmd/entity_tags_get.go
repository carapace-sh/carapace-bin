package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_tags_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the tags for a given entity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_tags_getCmd).Standalone()
	entity_tags_getCmd.Flags().StringP("guid", "g", "", "the entity GUID to retrieve tags for")
	entity_tagsCmd.AddCommand(entity_tags_getCmd)
}
