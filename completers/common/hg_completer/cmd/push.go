package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "push changes to the specified destination",
	GroupID: groups[group_remote_repository_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("all-bookmarks", false, "push all bookmarks (EXPERIMENTAL)")
	pushCmd.Flags().StringArrayP("bookmark", "B", nil, "bookmark to push")
	pushCmd.Flags().StringArrayP("branch", "b", nil, "a specific branch you would like to push")
	pushCmd.Flags().BoolP("force", "f", false, "force push")
	pushCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	pushCmd.Flags().Bool("new-branch", false, "allow pushing a new branch")
	pushCmd.Flags().Bool("publish", false, "push the changeset as public (EXPERIMENTAL)")
	pushCmd.Flags().String("pushvars", "", "variables that can be sent to server (ADVANCED)")
	pushCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	pushCmd.Flags().StringArrayP("rev", "r", nil, "a changeset intended to be included in the destination")
	pushCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).FlagCompletion(carapace.ActionMap{
		"bookmark":  hg.ActionBookmarks(),
		"branch":    hg.ActionBranches(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       hg.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(pushCmd).PositionalCompletion(
		carapace.Batch(hg.ActionPaths(), carapace.ActionFiles()).ToA(),
	)
}
