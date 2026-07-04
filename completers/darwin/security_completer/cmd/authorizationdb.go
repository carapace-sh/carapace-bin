package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authorizationdbCmd = &cobra.Command{
	Use:   "authorizationdb",
	Short: "Make changes to the authorization policy database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authorizationdbCmd).Standalone()
	rootCmd.AddCommand(authorizationdbCmd)
	carapace.Gen(authorizationdbCmd).PositionalCompletion(
		carapace.ActionValues("read", "write", "remove"),
	)
}
