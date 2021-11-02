package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var auth_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove authentication contexts ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_removeCmd).Standalone()
	authCmd.AddCommand(auth_removeCmd)
}
