package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var prPublishCmd = &cobra.Command{
	Use:     "pr-publish",
	Short:   "Publish bottles for a pull request with GitHub Actions",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prPublishCmd).Standalone()

	prPublishCmd.Flags().Bool("autosquash", false, "If supported on the target tap, automatically reformat and reword commits to our preferred format.")
	prPublishCmd.Flags().String("branch", "", "Branch to use the workflow from (default: `master`).")
	prPublishCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	prPublishCmd.Flags().BoolP("help", "h", false, "Show this message.")
	prPublishCmd.Flags().Bool("large-runner", false, "Run the upload job on a large runner.")
	prPublishCmd.Flags().String("message", "", "Message to include when autosquashing revision bumps, deletions, and rebuilds.")
	prPublishCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	prPublishCmd.Flags().String("tap", "", "Target tap repository (default: `homebrew/core`).")
	prPublishCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	prPublishCmd.Flags().String("workflow", "", "Target workflow filename (default: `publish-commit-bottles.yml`).")
	rootCmd.AddCommand(prPublishCmd)

	carapace.Gen(prPublishCmd).FlagCompletion(carapace.ActionMap{
		"tap": gh.ActionOwnerRepositories(gh.HostOpts{}),
	})
}
