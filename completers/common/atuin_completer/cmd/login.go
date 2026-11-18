package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().Bool("from-registration", false, "")
	loginCmd.Flags().BoolP("help", "h", false, "Print help")
	loginCmd.Flags().StringP("key", "k", "", "The encryption key for your account. Falls back to `ATUIN_ENCRYPTION_KEY`, then a prompt")
	loginCmd.Flags().StringP("password", "p", "", "Your password, or `-` to read it from stdin. Falls back to `ATUIN_PASSWORD`, then a prompt")
	loginCmd.Flags().StringP("totp-code", "t", "", "The two-factor authentication code for your account, if any")
	loginCmd.Flags().StringP("username", "u", "", "")
	loginCmd.Flag("from-registration").Hidden = true
	rootCmd.AddCommand(loginCmd)
}
