package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var labelsCmd = &cobra.Command{
	Use:     "labels",
	Short:   "Manage issue labels",
	GroupID: "ENTITIES",
	Aliases: []string{"label"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelsCmd).Standalone()

	labelsCmd.Flags().String("limit", "", "specify limit of items per page")
	labelsCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	labelsCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	labelsCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	labelsCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	labelsCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	labelsCmd.Flags().BoolP("save", "s", false, "Save all the labels as a file")
	rootCmd.AddCommand(labelsCmd)

	// TODO output formats
	carapace.Gen(labelsCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
