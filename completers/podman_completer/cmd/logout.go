package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout [options] [REGISTRY]",
	Short: "Log out of a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	logoutCmd.Flags().BoolP("all", "a", false, "Remove the cached credentials for all registries in the auth file")
	logoutCmd.Flags().String("authfile", "", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	logoutCmd.Flags().String("compat-auth-file", "", "path of a Docker-compatible config file to update instead")
	rootCmd.AddCommand(logoutCmd)
}
