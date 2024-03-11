package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Manage tags on New Relic entities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_tagsCmd).Standalone()
	entityCmd.AddCommand(entity_tagsCmd)
}
