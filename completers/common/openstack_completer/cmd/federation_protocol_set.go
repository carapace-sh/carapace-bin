package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_protocol_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set federation protocol properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_protocol_setCmd).Standalone()

	federation_protocol_setCmd.Flags().String("identity-provider", "", "Identity provider that supports <federation-protocol> (name or ID) (required)")
	federation_protocol_setCmd.Flags().String("mapping", "", "Mapping that is to be used (name or ID)")
	federation_protocol_setCmd.MarkFlagRequired("identity-provider")
	federation_protocolCmd.AddCommand(federation_protocol_setCmd)
}
