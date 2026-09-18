package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identity_provider_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete identity provider(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_provider_deleteCmd).Standalone()

	identity_providerCmd.AddCommand(identity_provider_deleteCmd)
}
