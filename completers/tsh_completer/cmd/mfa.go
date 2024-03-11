package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mfaCmd = &cobra.Command{
	Use:   "mfa",
	Short: "Manage multi-factor authentication (MFA) devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mfaCmd).Standalone()

	rootCmd.AddCommand(mfaCmd)
}
