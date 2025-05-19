package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/doing"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var workonCmd = &cobra.Command{
	Use:   "workon [OPTIONS] ISSUE",
	Short: "Create issue with PR and switch git branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workonCmd).Standalone()

	workonCmd.Flags().Bool("add-to-current-sprint", false, "Add item to the current sprint")
	workonCmd.Flags().Bool("auto-complete", false, "Set the PR to complete autom. when all policies have passed")
	workonCmd.Flags().String("branch-prefix", "", "The prefix to be prepended to the branch name")
	workonCmd.Flags().Bool("checkout", false, "Run git commands to checkout remote branch locally")
	workonCmd.Flags().StringP("default-branch", "b", "", "The name of the branch to branch from and to")
	workonCmd.Flags().Bool("delete-source-branch", false, "Set to delete source branch when pull request completes")
	workonCmd.Flags().Bool("do-not-add-to-current-sprint", false, "Do not add item to the current sprint")
	workonCmd.Flags().Bool("draft", false, "Create draft/WIP pull request")
	workonCmd.Flags().Bool("help", false, "Show this message and exit")
	workonCmd.Flags().StringP("label", "l", "", "Attach tags (labels) to work item")
	workonCmd.Flags().Bool("no-auto-complete", false, "Do not set the PR to complete autom. when all policies have passed")
	workonCmd.Flags().Bool("no-checkout", false, "Do not run git commands to checkout remote branch locally")
	workonCmd.Flags().Bool("no-delete-source-branch", false, "Set to delete source branch when pull request completes")
	workonCmd.Flags().Bool("no-draft", false, "Do not create draft/WIP pull request")
	workonCmd.Flags().Bool("no-self-approve", false, "Do not add yourself as reviewer and add your approval")
	workonCmd.Flags().StringP("parent", "p", "", "To create a child work item, specify the ID of the parent work item")
	workonCmd.Flags().StringP("reviewers", "r", "", "Space separated list of reviewer emails")
	workonCmd.Flags().Bool("self-approve", false, "Add yourself as reviewer and add your approval")
	workonCmd.Flags().StringP("story-points", "s", "", "The number of story points to assign")
	workonCmd.Flags().String("type", "", "Type of work item")
	rootCmd.AddCommand(workonCmd)

	carapace.Gen(workonCmd).FlagCompletion(carapace.ActionMap{
		"default-branch": git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}), // TODO test
		"label":          carapace.ActionValues(),                                                  // TODO
		"parent":         carapace.ActionValues(),                                                  // TODO
		"reviewers":      carapace.ActionValues(),                                                  // TODO
		"story-points":   carapace.ActionValues(),                                                  // TODO
		"type":           doing.ActionWorkItemTypes(),
	})

	// TODO positional completion
}
