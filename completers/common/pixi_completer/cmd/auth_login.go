package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store authentication information for a given host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().String("conda-token", "", "The token to use on anaconda.org / quetz authentication")
	auth_loginCmd.Flags().String("password", "", "The password to use (for basic HTTP authentication)")
	auth_loginCmd.Flags().String("s3-access-key-id", "", "The S3 access key ID")
	auth_loginCmd.Flags().String("s3-secret-access-key", "", "The S3 secret access key")
	auth_loginCmd.Flags().String("s3-session-token", "", "The S3 session token")
	auth_loginCmd.Flags().String("token", "", "The token to use (for authentication with prefix.dev)")
	auth_loginCmd.Flags().String("username", "", "The username to use (for basic HTTP authentication)")
	authCmd.AddCommand(auth_loginCmd)
}
