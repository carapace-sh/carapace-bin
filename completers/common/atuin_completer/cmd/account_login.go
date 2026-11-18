package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_loginCmd).Standalone()

	account_loginCmd.Flags().Bool("from-registration", false, "")
	account_loginCmd.Flags().BoolP("help", "h", false, "Print help")
	account_loginCmd.Flags().StringP("key", "k", "", "The encryption key for your account. Falls back to `ATUIN_ENCRYPTION_KEY`, then a prompt")
	account_loginCmd.Flags().StringP("password", "p", "", "Your password, or `-` to read it from stdin. Falls back to `ATUIN_PASSWORD`, then a prompt")
	account_loginCmd.Flags().StringP("totp-code", "t", "", "The two-factor authentication code for your account, if any")
	account_loginCmd.Flags().StringP("username", "u", "", "")
	account_loginCmd.Flag("from-registration").Hidden = true
	accountCmd.AddCommand(account_loginCmd)
}
