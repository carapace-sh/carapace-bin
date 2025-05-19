package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var labels_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a label",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labels_deleteCmd).Standalone()

	labels_deleteCmd.Flags().String("id", "", "label id")
	labels_deleteCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	labels_deleteCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	labels_deleteCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	labels_deleteCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	labelsCmd.AddCommand(labels_deleteCmd)

	// TODO completion
	carapace.Gen(labels_deleteCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(labels_deleteCmd).PositionalCompletion(
		action.ActionLabels(labels_deleteCmd),
	)
}
