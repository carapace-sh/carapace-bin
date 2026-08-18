package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a branch in the SVN repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().String("commit-url", "", "Use the specified URL to connect to the destination SVN repository")
	branchCmd.Flags().StringP("destination", "d", "", "Location of the branch in the SVN repository")
	branchCmd.Flags().Bool("dry-run", false, "Dry run")
	branchCmd.Flags().StringP("message", "m", "", "Commit message")
	branchCmd.Flags().Bool("parents", false, "Create parent folders")
	branchCmd.Flags().BoolP("tag", "t", false, "Create a tag instead of a branch")
	branchCmd.Flags().String("username", "", "SVN username to perform the commit as")
	rootCmd.AddCommand(branchCmd)

	carapace.Gen(branchCmd).FlagCompletion(carapace.ActionMap{
		"commit-url":  carapace.ActionValues(),
		"destination": carapace.ActionValues(),
		"message":     carapace.ActionValues(),
	})

	carapace.Gen(branchCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
