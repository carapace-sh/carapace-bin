package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incomingCmd = &cobra.Command{
	Use:     "incoming",
	Short:   "show new changesets found in source",
	Aliases: []string{"in"},
	GroupID: groups[group_remote_repository_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incomingCmd).Standalone()

	incomingCmd.Flags().BoolP("bookmarks", "B", false, "compare bookmarks")
	incomingCmd.Flags().StringArrayP("branch", "b", nil, "a specific branch you would like to pull")
	incomingCmd.Flags().String("bundle", "", "file to store the bundles into")
	incomingCmd.Flags().BoolP("force", "f", false, "run even if remote repository is unrelated")
	incomingCmd.Flags().BoolP("git", "g", false, "use git extended diff format")
	incomingCmd.Flags().BoolP("graph", "G", false, "show the revision DAG")
	incomingCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	incomingCmd.Flags().StringP("limit", "l", "", "limit number of changes displayed")
	incomingCmd.Flags().BoolP("newest-first", "n", false, "show newest record first")
	incomingCmd.Flags().BoolP("no-merges", "M", false, "do not show merges")
	incomingCmd.Flags().BoolP("patch", "p", false, "show patch")
	incomingCmd.Flags().Bool("remote-hidden", false, "include changesets hidden on the remote (EXPERIMENTAL)")
	incomingCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	incomingCmd.Flags().StringArrayP("rev", "r", nil, "a remote changeset intended to be added")
	incomingCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	incomingCmd.Flags().Bool("stat", false, "output diffstat-style summary of changes")
	incomingCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	incomingCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	incomingCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(incomingCmd)

	carapace.Gen(incomingCmd).FlagCompletion(carapace.ActionMap{
		"branch":    action.ActionBranches(),
		"bundle":    carapace.ActionFiles(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       action.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(incomingCmd).PositionalCompletion(
		carapace.Batch(action.ActionPaths(), carapace.ActionFiles()).ToA(),
	)
}
