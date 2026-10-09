package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var otp_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show OTP extension help",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_helpCmd).Standalone()
	otpCmd.AddCommand(otp_helpCmd)
}
