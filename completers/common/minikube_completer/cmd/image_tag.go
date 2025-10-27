package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_tagCmd = &cobra.Command{
	Use:     "tag",
	Short:   "Tag images",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_tagCmd).Standalone()

	imageCmd.AddCommand(image_tagCmd)
}
