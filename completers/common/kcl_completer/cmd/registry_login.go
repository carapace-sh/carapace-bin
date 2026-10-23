package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registry_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_loginCmd).Standalone()

	registry_loginCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls verification")
	registry_loginCmd.Flags().StringP("password", "p", "", "registry password or identity token")
	registry_loginCmd.Flags().Bool("password-stdin", false, "password from stdin")
	registry_loginCmd.Flags().StringP("username", "u", "", "registry username")
	registryCmd.AddCommand(registry_loginCmd)

	carapace.Gen(registry_loginCmd).PositionalCompletion(
		carapace.ActionValues("ghcr.io"),
	)
}
