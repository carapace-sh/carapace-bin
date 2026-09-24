package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete image(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_deleteCmd).Standalone()

	image_deleteCmd.Flags().String("store", "", "Store to delete image(s) from.")
	imageCmd.AddCommand(image_deleteCmd)
}
