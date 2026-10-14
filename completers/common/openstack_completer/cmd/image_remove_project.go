package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_remove_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Disassociate project with image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_remove_projectCmd).Standalone()

	image_remove_projectCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	image_removeCmd.AddCommand(image_remove_projectCmd)
}
