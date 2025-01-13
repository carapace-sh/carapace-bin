package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [OPTIONS] SERVICE [REPOSITORY[:TAG]]",
	Short: "Create a new image from a service container's changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().StringP("author", "a", "", "Author (e.g., \"John Hannibal Smith <hannibal@a-team.com>\")")
	commitCmd.Flags().StringP("change", "c", "", "Apply Dockerfile instruction to the created image")
	commitCmd.Flags().String("index", "", "index of the container if service has multiple replicas.")
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	commitCmd.Flags().BoolP("pause", "p", false, "Pause container during commit")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
