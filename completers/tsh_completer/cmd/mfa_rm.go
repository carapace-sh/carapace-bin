package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mfa_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a MFA device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mfa_rmCmd).Standalone()

	mfaCmd.AddCommand(mfa_rmCmd)
}
