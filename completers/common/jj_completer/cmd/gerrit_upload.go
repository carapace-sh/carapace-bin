package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var gerrit_uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload changes to Gerrit for code review, or update existing changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gerrit_uploadCmd).Standalone()

	gerrit_uploadCmd.Flags().StringSlice("cc", nil, "CC these emails on the change (can be repeated)")
	gerrit_uploadCmd.Flags().StringSlice("custom", nil, "Send the following custom keyed values to Gerrit (can be repeated)")
	gerrit_uploadCmd.Flags().String("deadline", "", "The deadline after which the push should be aborted")
	gerrit_uploadCmd.Flags().BoolP("dry-run", "n", false, "Only display what will change on the remote; do not push changes to Gerrit")
	gerrit_uploadCmd.Flags().Bool("edit", false, "Push the change as a change edit")
	gerrit_uploadCmd.Flags().StringSlice("hashtag", nil, "Apply a hashtag to the change (can be repeated)")
	gerrit_uploadCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gerrit_uploadCmd.Flags().Bool("ignore-attention-set", false, "Do not modify the attention set upon uploading")
	gerrit_uploadCmd.Flags().StringSliceP("label", "l", nil, "Add the following labels configured by Gerrit (can be repeated)")
	gerrit_uploadCmd.Flags().Bool("merged", false, "Create a new change, even if the change has already been merged")
	gerrit_uploadCmd.Flags().StringP("message", "m", "", "The description for the patch set")
	gerrit_uploadCmd.Flags().Bool("no-publish-comments", false, "Do not publish draft comments for the given change")
	gerrit_uploadCmd.Flags().String("notify", "", "Who to email notifications to (defaults to all)")
	gerrit_uploadCmd.Flags().StringSliceP("option", "o", nil, "Send a `git push -o` option (can be repeated)")
	gerrit_uploadCmd.Flags().Bool("private", false, "Mark the change as private")
	gerrit_uploadCmd.Flags().Bool("publish-comments", false, "Publish draft comments for the given change")
	gerrit_uploadCmd.Flags().Bool("ready", false, "Mark the change as ready (no longer work in progress)")
	gerrit_uploadCmd.Flags().String("remote", "", "The Gerrit remote to push to")
	gerrit_uploadCmd.Flags().StringP("remote-branch", "b", "", "The location where your changes are intended to land")
	gerrit_uploadCmd.Flags().Bool("remove-private", false, "Unmark the change as private")
	gerrit_uploadCmd.Flags().StringSlice("reviewer", nil, "Add these emails as a reviewer (can be repeated)")
	gerrit_uploadCmd.Flags().StringSlice("revision", nil, "The revisions to upload to Gerrit")
	gerrit_uploadCmd.Flags().Bool("skip-validation", false, "When --submit is provided, skip performing validations")
	gerrit_uploadCmd.Flags().Bool("submit", false, "Directly submit the changes, bypassing code review")
	gerrit_uploadCmd.Flags().String("topic", "", "Apply a topic to the change")
	gerrit_uploadCmd.Flags().String("trace", "", "For debugging Gerrit")
	gerrit_uploadCmd.Flags().Bool("wip", false, "Mark the change as WIP (work in progress)")
	gerrit_uploadCmd.Flag("revision").Hidden = true
	gerritCmd.AddCommand(gerrit_uploadCmd)

	carapace.Gen(gerrit_uploadCmd).FlagCompletion(carapace.ActionMap{
		"notify": carapace.ActionValuesDescribed(
			"none", "No emails",
			"owner", "Only the change owner is notified",
			"owner-reviewers", "Only the change owner and reviewers will be notified",
			"all", "All relevant users, including owner, reviewers, cc'd, users that have starred the change, and users who have configured a watch on files in the change",
		),
		"option":   git.ActionPushOptions(),
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()).UniqueList(","),
	})
}
