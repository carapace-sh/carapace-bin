package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webauthnwin_diagCmd = &cobra.Command{
	Use:    "diag",
	Short:  "Run windows webauthn diagnostics",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webauthnwin_diagCmd).Standalone()

	webauthnwinCmd.AddCommand(webauthnwin_diagCmd)
}
