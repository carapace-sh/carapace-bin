package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_deleteCmd = &cobra.Command{
	Use:   "delete [token-value] ...",
	Short: "Delete bootstrap tokens on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_deleteCmd).Standalone()

	tokenCmd.AddCommand(token_deleteCmd)
}
