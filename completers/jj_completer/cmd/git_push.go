package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push to a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_pushCmd).Standalone()

	git_pushCmd.Flags().Bool("all", false, "Push all bookmarks (including deleted bookmarks)")
	git_pushCmd.Flags().Bool("allow-empty-description", false, "Allow commits with empty description messages to be pushed")
	git_pushCmd.Flags().BoolP("allow-new", "N", false, "Allow pushing new bookmarks")
	git_pushCmd.Flags().Bool("allow-private", false, "Allow pushing commits that are private")
	git_pushCmd.Flags().StringSliceP("bookmark", "b", nil, "Push only this bookmark (can be repeated)")
	git_pushCmd.Flags().StringSliceP("change", "c", nil, "Push this commit by creating a bookmark based on its change ID (can be repeated)")
	git_pushCmd.Flags().Bool("deleted", false, "Push all deleted bookmarks")
	git_pushCmd.Flags().Bool("dry-run", false, "Only display what will change on the remote")
	git_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_pushCmd.Flags().String("remote", "", "The remote to push to (only named remotes are supported)")
	git_pushCmd.Flags().StringSliceP("revisions", "r", nil, "Push bookmarks pointing to these commits")
	gitCmd.AddCommand(git_pushCmd)

	carapace.Gen(git_pushCmd).FlagCompletion(carapace.ActionMap{
		"bookmark":  jj.ActionLocalBookmarks(),
		"change":    carapace.ActionValues(), // TODO
		"remote":    jj.ActionRemotes(),
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
