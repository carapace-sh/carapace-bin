package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var otp_codeCmd = &cobra.Command{
	Use:   "code",
	Short: "generate and print an OTP code from the secret key stored",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_codeCmd).Standalone()
	otp_codeCmd.Flags().BoolP("clip", "c", false, "copy to clipboard")

	otpCmd.AddCommand(otp_codeCmd)

	carapace.Gen(otp_codeCmd).PositionalCompletion(
		pass.ActionPasswords(),
	)
}
