package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var logins_editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit Gitea logins",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logins_editCmd).Standalone()

	logins_editCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	loginsCmd.AddCommand(logins_editCmd)

	// TODO completion
	carapace.Gen(logins_editCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
	})
}
