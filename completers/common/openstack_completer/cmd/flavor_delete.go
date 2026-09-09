package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flavor_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete flavor(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavor_deleteCmd).Standalone()

	flavorCmd.AddCommand(flavor_deleteCmd)
}
