package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pr_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] WORK_ITEM_ID",
	Short: "Create a pull request from a work item ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_createCmd).Standalone()

	pr_createCmd.Flags().Bool("auto-complete", false, "Set the PR to complete autom. when all policies have passed")
	pr_createCmd.Flags().String("branch-prefix", "", "The prefix to be prepended to the branch name")
	pr_createCmd.Flags().Bool("checkout", false, "Run git commands to checkout remote branch locally")
	pr_createCmd.Flags().StringP("default-branch", "b", "", "The name of the branch to branch from and to")
	pr_createCmd.Flags().Bool("delete-source-branch", false, "Set to delete source branch when pull request completes")
	pr_createCmd.Flags().Bool("draft", false, "Create draft/WIP pull request")
	pr_createCmd.Flags().Bool("help", false, "Show this message and exit")
	pr_createCmd.Flags().Bool("no-auto-complete", false, "Set the PR to complete autom. when all policies have passed")
	pr_createCmd.Flags().Bool("no-checkout", false, "Do not run git commands to checkout remote branch locally")
	pr_createCmd.Flags().Bool("no-delete-source-branch", false, "Set to delete source branch when pull request completes")
	pr_createCmd.Flags().Bool("no-draft", false, "Do not create draft/WIP pull request")
	pr_createCmd.Flags().Bool("no-self-approve", false, "Do not add yourself as reviewer and add your approval")
	pr_createCmd.Flags().Bool("no-web", false, "Open newly created issue in the web browser")
	pr_createCmd.Flags().StringP("reviewers", "r", "", "Space separated list of reviewer emails or aliases")
	pr_createCmd.Flags().Bool("self-approve", false, "Add yourself as reviewer and add your approval")
	pr_createCmd.Flags().BoolP("web", "w", false, "Open newly created issue in the web browser")
	prCmd.AddCommand(pr_createCmd)

	carapace.Gen(pr_createCmd).FlagCompletion(carapace.ActionMap{
		"branch-prefix":  carapace.ActionValues(),
		"default-branch": git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}), // TODO test
		"reviewers":      carapace.ActionValues(),
	})

	// TODO positional completion
}
