package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fido2_diagCmd = &cobra.Command{
	Use:    "diag",
	Short:  "Run FIDO2 diagnostics",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fido2_diagCmd).Standalone()

	fido2Cmd.AddCommand(fido2_diagCmd)
}
