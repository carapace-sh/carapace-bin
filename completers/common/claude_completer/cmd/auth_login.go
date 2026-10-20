package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Sign in to your Anthropic account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().Bool("claudeai", false, "Use Claude subscription (default)")
	auth_loginCmd.Flags().Bool("console", false, "Use Anthropic Console (API usage billing) instead of Claude subscription")
	auth_loginCmd.Flags().String("email", "", "Pre-populate email address on the login page")
	auth_loginCmd.Flags().BoolP("help", "h", false, "Display help for command")
	auth_loginCmd.Flags().Bool("sso", false, "Force SSO login flow")
	authCmd.AddCommand(auth_loginCmd)
}
