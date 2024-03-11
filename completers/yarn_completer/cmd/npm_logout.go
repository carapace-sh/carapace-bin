package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the npm registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_logoutCmd).Standalone()

	npm_logoutCmd.Flags().BoolP("all", "A", false, "Logout of all registries")
	npm_logoutCmd.Flags().Bool("publish", false, "Logout of the publish registry")
	npm_logoutCmd.Flags().StringP("scope", "s", "", "Logout of the registry configured for a given scope")
	npmCmd.AddCommand(npm_logoutCmd)

	// TODO scope
}
