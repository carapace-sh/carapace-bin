package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var labels_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List labels",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labels_listCmd).Standalone()

	labels_listCmd.Flags().String("limit", "", "specify limit of items per page")
	labels_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	labels_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	labels_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	labels_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	labels_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	labels_listCmd.Flags().BoolP("save", "s", false, "Save all the labels as a file")
	labelsCmd.AddCommand(labels_listCmd)

	// TODO completion
	carapace.Gen(labels_listCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
