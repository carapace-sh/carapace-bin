package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in Docker to a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_loginCmd).Standalone()
	registry_loginCmd.Flags().Int("expiry-seconds", 0, "The length of time the registry credentials will be valid for in seconds. By default, the credentials do not expire.")
	registryCmd.AddCommand(registry_loginCmd)
}
