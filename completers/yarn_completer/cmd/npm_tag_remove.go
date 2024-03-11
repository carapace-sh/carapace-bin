package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_tag_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a tag from a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_tag_removeCmd).Standalone()

	npm_tagCmd.AddCommand(npm_tag_removeCmd)

	// TODO positional completion
}
