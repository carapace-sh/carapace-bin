package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ultrareviewCmd = &cobra.Command{
	Use:   "ultrareview",
	Short: "Run a cloud-hosted multi-agent code review of the current branch (or a PR number / base branch)…",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ultrareviewCmd).Standalone()

	ultrareviewCmd.Flags().BoolP("help", "h", false, "Display help for command")
	ultrareviewCmd.Flags().Bool("json", false, "Print the raw bugs.json payload instead of formatted findings")
	ultrareviewCmd.Flags().Bool("no-post", false, "Do not post the findings to the PR…")
	ultrareviewCmd.Flags().Bool("post", false, "Post the finished review's findings to the PR as you…")
	ultrareviewCmd.Flags().String("timeout", "", "Maximum minutes to wait for the review to finish")
	rootCmd.AddCommand(ultrareviewCmd)
}
