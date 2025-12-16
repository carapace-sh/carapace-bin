package cmd

import (
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pass",
	Short: "stores, retrieves, generates, and synchronizes passwords securely",
	Long:  "https://www.passwordstore.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		pass.ActionPasswords(),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		otpCmd.Hidden = true

		// TODO support additional installed extension(s)
		potentialDirs := []string{"/usr/lib/password-store/extensions/otp.bash"}
		if nixPofile, err := traverse.NixProfile(carapace.NewContext()); err == nil { // Support home-manager
			potentialDirs = append(potentialDirs, filepath.Join(nixPofile, "/lib/password-store/extensions/otp.bash"))
		}

		for _, path := range potentialDirs {
			if _, err := os.Stat(path); err != nil {
				otpCmd.Hidden = false
				return
			}
		}
	})
}
