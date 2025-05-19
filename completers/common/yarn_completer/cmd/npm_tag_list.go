package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_tag_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all dist-tags of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_tag_listCmd).Standalone()

	npm_tagCmd.AddCommand(npm_tag_listCmd)

	// TODO positional completion
}
