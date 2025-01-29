package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_existsCmd = &cobra.Command{
	Use:   "exists IMAGE",
	Short: "Check if an image exists in local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_existsCmd).Standalone()

	imageCmd.AddCommand(image_existsCmd)
}
