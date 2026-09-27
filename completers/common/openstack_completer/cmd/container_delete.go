package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_deleteCmd).Standalone()

	container_deleteCmd.Flags().BoolP("recursive", "r", false, "Recursively delete objects and container")
	containerCmd.AddCommand(container_deleteCmd)
}
