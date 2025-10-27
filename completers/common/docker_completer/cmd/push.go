package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push [OPTIONS] NAME[:TAG]",
	Short:   "Upload an image to a registry",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().BoolP("all-tags", "a", false, "Push all tags of an image to the repository")
	pushCmd.Flags().Bool("disable-content-trust", false, "Skip image signing")
	pushCmd.Flags().String("platform", "", "Push a platform-specific manifest as a single-platform image to the registry.")
	pushCmd.Flags().BoolP("quiet", "q", false, "Suppress verbose output")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
