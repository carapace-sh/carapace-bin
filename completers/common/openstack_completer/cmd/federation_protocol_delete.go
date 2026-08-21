package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_protocol_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete federation protocol(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_protocol_deleteCmd).Standalone()

	federation_protocol_deleteCmd.Flags().String("identity-provider", "", "Identity provider that supports <federation-protocol> (name or ID) (required)")
	federation_protocol_deleteCmd.MarkFlagRequired("identity-provider")
	federation_protocolCmd.AddCommand(federation_protocol_deleteCmd)
}
