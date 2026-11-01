package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _comment_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a comment anchored to a line in a diff",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_comment_addCmd).Standalone()

	_comment_addCmd.Flags().String("commit", "", "Anchor to the diff of this commit instead of the uncommitted changes")
	_comment_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	_comment_addCmd.Flags().StringP("message", "m", "", "The comment text")
	_comment_addCmd.Flags().Bool("old", false, "Count the line in the old side of the diff (removed lines and context) instead of the new side")
	_comment_addCmd.MarkFlagRequired("message")
	_commentCmd.AddCommand(_comment_addCmd)
}
