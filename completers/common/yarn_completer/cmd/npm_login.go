package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store new login info to access the npm registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_loginCmd).Standalone()

	npm_loginCmd.Flags().Bool("publish", false, "Login to the publish registry")
	npm_loginCmd.Flags().StringP("scope", "s", "", "Login to the registry configured for a given scope")
	npmCmd.AddCommand(npm_loginCmd)

	// TODO scope
}
