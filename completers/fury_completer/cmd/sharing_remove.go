package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharing_removeCmd = &cobra.Command{
	Use:   "remove EMAIL",
	Short: "Remove a collaborator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharing_removeCmd).Standalone()

	sharingCmd.AddCommand(sharing_removeCmd)
}
