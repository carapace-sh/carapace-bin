package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_dockerHelperCmd = &cobra.Command{
	Use:   "docker-helper",
	Short: "A Docker credential helper for GitLab container registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_dockerHelperCmd).Standalone()

	authCmd.AddCommand(auth_dockerHelperCmd)
}
