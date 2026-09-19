package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outgoingCmd = &cobra.Command{
	Use:     "outgoing",
	Short:   "show changesets not found in the destination",
	Aliases: []string{"out"},
	GroupID: groups[group_remote_repository_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outgoingCmd).Standalone()

	outgoingCmd.Flags().BoolP("bookmarks", "B", false, "compare bookmarks")
	outgoingCmd.Flags().StringArrayP("branch", "b", nil, "a specific branch you would like to push")
	outgoingCmd.Flags().BoolP("force", "f", false, "run even when the destination is unrelated")
	outgoingCmd.Flags().BoolP("git", "g", false, "use git extended diff format")
	outgoingCmd.Flags().BoolP("graph", "G", false, "show the revision DAG")
	outgoingCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	outgoingCmd.Flags().StringP("limit", "l", "", "limit number of changes displayed")
	outgoingCmd.Flags().BoolP("newest-first", "n", false, "show newest record first")
	outgoingCmd.Flags().BoolP("no-merges", "M", false, "do not show merges")
	outgoingCmd.Flags().BoolP("patch", "p", false, "show patch")
	outgoingCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	outgoingCmd.Flags().StringArrayP("rev", "r", nil, "a changeset intended to be included in the destination")
	outgoingCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	outgoingCmd.Flags().Bool("stat", false, "output diffstat-style summary of changes")
	outgoingCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	outgoingCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	outgoingCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(outgoingCmd)

	carapace.Gen(outgoingCmd).FlagCompletion(carapace.ActionMap{
		"branch":    action.ActionBranches(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       action.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(outgoingCmd).PositionalCompletion(
		carapace.Batch(action.ActionPaths(), carapace.ActionFiles()).ToA(),
	)
}
