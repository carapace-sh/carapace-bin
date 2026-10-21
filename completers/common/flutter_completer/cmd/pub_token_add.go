package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_token_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an authentication token for a package repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_token_addCmd).Standalone()

	pub_token_addCmd.Flags().String("env-var", "", "Read the secret token from this environment variable when making requests.")
	pub_token_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_tokenCmd.AddCommand(pub_token_addCmd)
}
