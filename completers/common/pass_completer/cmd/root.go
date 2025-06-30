package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pass",
	Short: "stores, retrieves, generates, and synchronizes passwords securely",
	Long:  "https://www.passwordstore.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		pass.ActionPasswords(),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		otpCmd.Hidden = true

		homeDir, err := os.UserHomeDir()
		if err != nil {
			// Should never happen...
			return
		}

		// TODO support additional installed extension(s)
		potentialDirs := []string{
			"/usr/lib/password-store/extensions/otp.bash",
			// Support home-manager
			homeDir + "/.nix-profile/lib/password-store/extensions/otp.bash",
		}
		for _, path := range potentialDirs {
			if _, err := os.Stat(path); err != nil {
				otpCmd.Hidden = false
				return
			}
		}
	})
}
