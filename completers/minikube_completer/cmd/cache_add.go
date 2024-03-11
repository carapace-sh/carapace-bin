package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an image to local cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_addCmd).Standalone()
	cache_addCmd.Flags().Bool("", false, "Add image to cache for all running minikube clusters")
	cacheCmd.AddCommand(cache_addCmd)

	carapace.Gen(cache_addCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
