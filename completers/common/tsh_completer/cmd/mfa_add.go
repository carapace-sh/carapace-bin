package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mfa_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new MFA device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mfa_addCmd).Standalone()

	mfa_addCmd.Flags().String("name", "", "Name of the new MFA device")
	mfa_addCmd.Flags().String("type", "", "Type of the new MFA device (TOTP, WEBAUTHN)")
	mfaCmd.AddCommand(mfa_addCmd)

	carapace.Gen(mfa_addCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("TOTP", "WEBAUTHN"),
	})
}
