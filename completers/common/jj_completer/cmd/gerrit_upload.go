package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
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
	gerrit_uploadCmd.Flags().BoolP("dry-run", "n", false, "Do not actually push the changes to Gerrit")
	gerrit_uploadCmd.Flags().String("hashtag", "", "Applies a hashtag to the change")
	gerrit_uploadCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gerrit_uploadCmd.Flags().Bool("ignore-attention-set", false, "Do not modify the attention set upon uploading")
	gerrit_uploadCmd.Flags().StringSliceP("label", "l", nil, "Add the following labels configured by Gerrit (can be repeated)")
	gerrit_uploadCmd.Flags().Bool("no-publish-comments", false, "Disables publishing of any draft comments for the given change")
	gerrit_uploadCmd.Flags().String("notify", "", "Who to email notifications to (defaults to all)")
	gerrit_uploadCmd.Flags().Bool("private", false, "Marks the change as private")
	gerrit_uploadCmd.Flags().Bool("publish-comments", false, "Publishes any draft comments for the given change")
	gerrit_uploadCmd.Flags().Bool("ready", false, "Unmarks the change as WIP (work in progress)")
	gerrit_uploadCmd.Flags().String("remote", "", "The Gerrit remote to push to")
	gerrit_uploadCmd.Flags().StringP("remote-branch", "b", "", "The location where your changes are intended to land")
	gerrit_uploadCmd.Flags().Bool("remove-private", false, "Unmarks the change as private")
	gerrit_uploadCmd.Flags().StringSlice("reviewer", nil, "Add these emails as a reviewer (can be repeated)")
	gerrit_uploadCmd.Flags().StringSliceP("revision", "r", nil, "The revset, selecting which revisions are sent in to Gerrit")
	gerrit_uploadCmd.Flags().StringSlice("revisions", nil, "The revset, selecting which revisions are sent in to Gerrit")
	gerrit_uploadCmd.Flags().Bool("skip-validation", false, "When --submit is provided, skip performing validations")
	gerrit_uploadCmd.Flags().Bool("submit", false, "Directly submit the changes, bypassing code review")
	gerrit_uploadCmd.Flags().String("topic", "", "Applies a topic to the change")
	gerrit_uploadCmd.Flags().String("trace", "", "For debugging Gerrit")
	gerrit_uploadCmd.Flags().Bool("wip", false, "Marks the change as WIP (work in progress)")
	gerrit_uploadCmd.Flag("revisions").Hidden = true
	gerritCmd.AddCommand(gerrit_uploadCmd)

	// TODO complete gerrit flags
	carapace.Gen(gerrit_uploadCmd).FlagCompletion(carapace.ActionMap{
		"revision":  jj.ActionRevs(jj.RevOption{}.Default()).UniqueList(","),
		"revisions": jj.ActionRevs(jj.RevOption{}.Default()).UniqueList(","),
	})
}
