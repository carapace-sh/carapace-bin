package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a tag in the SVN repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()

	tagCmd.Flags().String("commit-url", "", "Use the specified URL to connect to the destination SVN repository")
	tagCmd.Flags().StringP("destination", "d", "", "Location of the tag in the SVN repository")
	tagCmd.Flags().BoolP("dry-run", "n", false, "Dry run")
	tagCmd.Flags().StringP("message", "m", "", "Commit message")
	tagCmd.Flags().Bool("parents", false, "Create parent folders")
	tagCmd.Flags().String("username", "", "SVN username to perform the commit as")
	rootCmd.AddCommand(tagCmd)

	carapace.Gen(tagCmd).FlagCompletion(carapace.ActionMap{
		"commit-url":  carapace.ActionValues(),
		"destination": carapace.ActionValues(),
		"message":     carapace.ActionValues(),
	})

	carapace.Gen(tagCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
