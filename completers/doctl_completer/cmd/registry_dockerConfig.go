package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_dockerConfigCmd = &cobra.Command{
	Use:   "docker-config",
	Short: "Generate a docker auth configuration for a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_dockerConfigCmd).Standalone()
	registry_dockerConfigCmd.Flags().Int("expiry-seconds", 0, "The length of time the registry credentials will be valid for in seconds. By default, the credentials do not expire.")
	registry_dockerConfigCmd.Flags().Bool("read-write", false, "Generate credentials that can push to your registry")
	registryCmd.AddCommand(registry_dockerConfigCmd)
}
