package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registered_limit_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a registered limit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registered_limit_deleteCmd).Standalone()

	registered_limitCmd.AddCommand(registered_limit_deleteCmd)
}
