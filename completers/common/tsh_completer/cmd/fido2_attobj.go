package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fido2_attobjCmd = &cobra.Command{
	Use:    "attobj",
	Short:  "Parse a stored attestation object",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fido2_attobjCmd).Standalone()

	fido2Cmd.AddCommand(fido2_attobjCmd)
}
