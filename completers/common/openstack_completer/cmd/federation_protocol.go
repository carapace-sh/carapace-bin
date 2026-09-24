package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_protocolCmd = &cobra.Command{
	Use:   "protocol",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_protocolCmd).Standalone()

	federationCmd.AddCommand(federation_protocolCmd)
}
