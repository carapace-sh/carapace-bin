package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var otp_uriCmd = &cobra.Command{
	Use:   "uri",
	Short: "Print the key URI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_uriCmd).Standalone()
	otp_uriCmd.Flags().BoolP("clip", "c", false, "copy to clipboard")
	otp_uriCmd.Flags().BoolP("qrcode", "q", false, "display qr code using qrencode")

	otpCmd.AddCommand(otp_uriCmd)

	carapace.Gen(otp_uriCmd).PositionalCompletion(
		pass.ActionPasswords(),
	)
}
