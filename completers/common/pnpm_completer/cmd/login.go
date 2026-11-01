package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Log in to an npm registry",
	Aliases: []string{"adduser"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	loginCmd.Flags().String("registry", "", "The registry to log in to")
	loginCmd.Flags().String("scope", "", "Associate the login token with a package scope and record the scope-to-registry mapping")
	rootCmd.AddCommand(loginCmd)
}
