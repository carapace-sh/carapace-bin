package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webauthnwinCmd = &cobra.Command{
	Use:    "webauthnwin",
	Short:  "Manage Windows WebAuthn",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webauthnwinCmd).Standalone()

	rootCmd.AddCommand(webauthnwinCmd)
}
