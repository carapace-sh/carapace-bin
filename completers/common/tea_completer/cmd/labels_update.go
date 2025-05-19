package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var labels_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a label",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labels_updateCmd).Standalone()

	labels_updateCmd.Flags().String("color", "", "label color value")
	labels_updateCmd.Flags().String("description", "", "label description")
	labels_updateCmd.Flags().String("id", "", "label id")
	labels_updateCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	labels_updateCmd.Flags().String("name", "", "label name")
	labels_updateCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	labels_updateCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	labels_updateCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	labelsCmd.AddCommand(labels_updateCmd)

	// TODO completion
	carapace.Gen(labels_updateCmd).FlagCompletion(carapace.ActionMap{
		"color":  color.ActionHexColors(),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(labels_updateCmd).PositionalCompletion(
		action.ActionLabels(labels_updateCmd),
	)
}
