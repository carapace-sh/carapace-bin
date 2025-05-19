package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var passportClientCmd = &cobra.Command{
	Use:   "passport:client [--personal] [--password] [--client] [--name [NAME]] [--provider [PROVIDER]] [--redirect_uri [REDIRECT_URI]] [--user_id [USER_ID]] [--public]",
	Short: "Create a client for issuing access tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(passportClientCmd).Standalone()

	passportClientCmd.Flags().Bool("client", false, "Create a client credentials grant client")
	passportClientCmd.Flags().String("name", "", "The name of the client")
	passportClientCmd.Flags().Bool("password", false, "Create a password grant client")
	passportClientCmd.Flags().Bool("personal", false, "Create a personal access token client")
	passportClientCmd.Flags().String("provider", "", "The name of the user provider")
	passportClientCmd.Flags().Bool("public", false, "Create a public client (Auth code grant type only)")
	passportClientCmd.Flags().String("redirect_uri", "", "The URI to redirect to after authorization")
	passportClientCmd.Flags().String("user_id", "", "The user ID the client should be assigned to")
	rootCmd.AddCommand(passportClientCmd)
}
