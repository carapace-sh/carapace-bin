package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset image tags and properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_unsetCmd).Standalone()

	image_unsetCmd.Flags().String("property", "", "Unset a property on this image (repeat option to unset multiple properties)")
	image_unsetCmd.Flags().String("tag", "", "Unset a tag on this image (repeat option to unset multiple tags)")
	imageCmd.AddCommand(image_unsetCmd)
}
