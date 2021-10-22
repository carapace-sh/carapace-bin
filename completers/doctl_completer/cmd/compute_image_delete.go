package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_image_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete an image from your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_deleteCmd).Standalone()
	compute_image_deleteCmd.Flags().BoolP("force", "f", false, "Force image delete")
	compute_imageCmd.AddCommand(compute_image_deleteCmd)
}
