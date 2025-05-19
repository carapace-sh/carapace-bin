package cmd

import (
	"github.com/carapace-sh/carapace"
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
	prPublishCmd.Flags().Bool("branch", false, "Branch to use the workflow from (default: `master`).")
	prPublishCmd.Flags().Bool("debug", false, "Display any debugging information.")
	prPublishCmd.Flags().Bool("help", false, "Show this message.")
	prPublishCmd.Flags().Bool("large-runner", false, "Run the upload job on a large runner.")
	prPublishCmd.Flags().Bool("message", false, "Message to include when autosquashing revision bumps, deletions, and rebuilds.")
	prPublishCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	prPublishCmd.Flags().Bool("tap", false, "Target tap repository (default: `homebrew/core`).")
	prPublishCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	prPublishCmd.Flags().Bool("workflow", false, "Target workflow filename (default: `publish-commit-bottles.yml`).")
	rootCmd.AddCommand(prPublishCmd)
}
