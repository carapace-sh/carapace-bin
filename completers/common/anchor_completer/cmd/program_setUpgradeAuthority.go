package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var program_setUpgradeAuthorityCmd = &cobra.Command{
	Use:   "set-upgrade-authority",
	Short: "Set a new program authority",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_setUpgradeAuthorityCmd).Standalone()

	program_setUpgradeAuthorityCmd.Flags().Bool("final", false, "Make the program immutable (cannot be upgraded)")
	program_setUpgradeAuthorityCmd.Flags().BoolP("help", "h", false, "Print help")
	program_setUpgradeAuthorityCmd.Flags().String("new-upgrade-authority", "", "New upgrade authority pubkey")
	program_setUpgradeAuthorityCmd.Flags().String("new-upgrade-authority-signer", "", "New upgrade authority signer (keypair file). Required unless --skip-new-upgrade-authority-signer-check is used. When provided, both current and new authority will sign (checked mode, recommended)")
	program_setUpgradeAuthorityCmd.Flags().Bool("skip-new-upgrade-authority-signer-check", false, "Skip new upgrade authority signer check. Allows setting authority with only current authority signature. WARNING: Less safe - use only if you're confident the pubkey is correct")
	program_setUpgradeAuthorityCmd.Flags().String("upgrade-authority", "", "Current upgrade authority keypair (defaults to configured wallet)")
	programCmd.AddCommand(program_setUpgradeAuthorityCmd)
}
