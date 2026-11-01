package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registry_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_logoutCmd).Standalone()

	registryCmd.AddCommand(registry_logoutCmd)

	carapace.Gen(registry_logoutCmd).PositionalCompletion(
		carapace.ActionValues("ghcr.io"),
	)
}
