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

	loginCmd.Flags().BoolP("help", "h", false, "Print help")
	loginCmd.Flags().StringP("key", "k", "", "The encryption key for your account")
	loginCmd.Flags().StringP("password", "p", "", "")
	loginCmd.Flags().StringP("username", "u", "", "")
	rootCmd.AddCommand(loginCmd)
}
