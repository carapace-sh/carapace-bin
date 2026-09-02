package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_rmCmd = &cobra.Command{
	Use:     "rm",
	Short:   "remove an authentication token from the registry",
	Aliases: []string{"delete", "revoke", "remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_rmCmd).Standalone()
	tokenCmd.AddCommand(token_rmCmd)
}
