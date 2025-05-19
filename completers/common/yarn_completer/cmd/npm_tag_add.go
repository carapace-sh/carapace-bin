package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_tag_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a tag for a specific version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_tag_addCmd).Standalone()

	npm_tagCmd.AddCommand(npm_tag_addCmd)

	// TODO positional completion
}
