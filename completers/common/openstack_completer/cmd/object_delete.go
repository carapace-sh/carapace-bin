package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete object from container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_deleteCmd).Standalone()

	objectCmd.AddCommand(object_deleteCmd)
}
