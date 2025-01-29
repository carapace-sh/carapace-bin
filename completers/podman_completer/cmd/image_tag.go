package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_tagCmd = &cobra.Command{
	Use:   "tag IMAGE TARGET_NAME [TARGET_NAME...]",
	Short: "Add an additional name to a local image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_tagCmd).Standalone()

	imageCmd.AddCommand(image_tagCmd)
}
