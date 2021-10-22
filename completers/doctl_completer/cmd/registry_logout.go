package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out Docker from a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_logoutCmd).Standalone()
	registryCmd.AddCommand(registry_logoutCmd)
}
