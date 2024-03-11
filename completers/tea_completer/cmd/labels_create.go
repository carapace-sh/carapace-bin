package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var labels_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a label",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labels_createCmd).Standalone()

	labels_createCmd.Flags().String("color", "", "label color value")
	labels_createCmd.Flags().String("description", "", "label description")
	labels_createCmd.Flags().String("file", "", "indicate a label file")
	labels_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	labels_createCmd.Flags().String("name", "", "label name")
	labels_createCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	labels_createCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	labels_createCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	labelsCmd.AddCommand(labels_createCmd)

	// TODO completion
	carapace.Gen(labels_createCmd).FlagCompletion(carapace.ActionMap{
		"color":  color.ActionHexColors(),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
