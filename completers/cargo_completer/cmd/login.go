package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save an api token from the registry locally. If token is not specified, it will be read from stdin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().Bool("generate-keypair", false, "Generate a public/secret keypair (unstable)")
	loginCmd.Flags().BoolP("help", "h", false, "Print help")
	loginCmd.Flags().String("key-subject", "", "Set the key subject for this registry (unstable)")
	loginCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	loginCmd.Flags().String("registry", "", "Registry to use")
	loginCmd.Flags().Bool("secret-key", false, "Prompt for secret key (unstable)")
	rootCmd.AddCommand(loginCmd)
}
