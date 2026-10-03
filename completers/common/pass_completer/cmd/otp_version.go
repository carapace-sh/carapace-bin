package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var otp_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show OTP extension version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otp_versionCmd).Standalone()
	otpCmd.AddCommand(otp_versionCmd)
}
