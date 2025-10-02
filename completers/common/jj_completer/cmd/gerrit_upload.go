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

	gerrit_uploadCmd.Flags().BoolP("dry-run", "n", false, "Do not actually push the changes to Gerrit")
	gerrit_uploadCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gerrit_uploadCmd.Flags().String("remote", "", "The Gerrit remote to push to")
	gerrit_uploadCmd.Flags().StringP("remote-branch", "b", "", "The location where your changes are intended to land")
	gerrit_uploadCmd.Flags().StringSliceP("revisions", "r", nil, "The revset, selecting which revisions are sent in to Gerrit")
	gerritCmd.AddCommand(gerrit_uploadCmd)

	// TODO complete gerrit flags
	carapace.Gen(gerrit_uploadCmd).FlagCompletion(carapace.ActionMap{
		"revisions": jj.ActionRevs(jj.RevOption{}.Default()).UniqueList(","),
	})
}
