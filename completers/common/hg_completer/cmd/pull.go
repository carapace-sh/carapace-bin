package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "pull changes from the specified source",
	GroupID: groups[group_remote_repository_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().StringArrayP("bookmark", "B", nil, "bookmark to pull")
	pullCmd.Flags().StringArrayP("branch", "b", nil, "a specific branch you would like to pull")
	pullCmd.Flags().Bool("confirm", false, "confirm pull before applying changes")
	pullCmd.Flags().BoolP("force", "f", false, "run even when remote repository is unrelated")
	pullCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	pullCmd.Flags().Bool("remote-hidden", false, "include changesets hidden on the remote (EXPERIMENTAL)")
	pullCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	pullCmd.Flags().StringArrayP("rev", "r", nil, "a remote changeset intended to be added")
	pullCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	pullCmd.Flags().BoolP("update", "u", false, "update to new branch head if new descendants were pulled")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"bookmark":  action.ActionBookmarks(),
		"branch":    action.ActionBranches(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       action.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(pullCmd).PositionalCompletion(
		carapace.Batch(action.ActionPaths(), carapace.ActionFiles()).ToA(),
	)
}
