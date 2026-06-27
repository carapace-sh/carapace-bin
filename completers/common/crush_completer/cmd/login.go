package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/crush"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login [platform]",
	Short:   "Login Crush to a platform",
	Aliases: []string{"auth"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().BoolP("force", "f", false, "Force re-authentication even if already logged in")
	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).PositionalCompletion(
		crush.ActionPlatforms(),
	)
}
