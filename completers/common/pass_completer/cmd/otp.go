package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var otpCmd = &cobra.Command{
	Use:   "otp",
	Short: "extension for managing one-time-password (OTP) tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(otpCmd).Standalone()
	otpCmd.Flags().BoolP("clip", "c", false, "Copy the OTP code to the clipboard")
	otpCmd.Flags().BoolP("help", "h", false, "Show help")
	otpCmd.Flags().BoolP("quiet", "q", false, "Print only the OTP code")
	otpCmd.Flags().Bool("version", false, "Show version information")

	rootCmd.AddCommand(otpCmd)

	carapace.Gen(otpCmd).PositionalCompletion(pass.ActionPasswords())
}
