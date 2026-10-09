package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identity_provider_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set identity provider properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_provider_setCmd).Standalone()

	identity_provider_setCmd.Flags().String("authorization-ttl", "", "Time to keep the role assignments for users authenticating via this identity provider.")
	identity_provider_setCmd.Flags().String("description", "", "Set identity provider description")
	identity_provider_setCmd.Flags().Bool("disable", false, "Disable the identity provider")
	identity_provider_setCmd.Flags().Bool("enable", false, "Enable the identity provider")
	identity_provider_setCmd.Flags().String("remote-id", "", "Remote IDs to associate with the Identity Provider (repeat option to provide multiple values)")
	identity_provider_setCmd.Flags().String("remote-id-file", "", "Name of a file that contains many remote IDs to associate with the identity provider, one per line")
	identity_providerCmd.AddCommand(identity_provider_setCmd)
}
