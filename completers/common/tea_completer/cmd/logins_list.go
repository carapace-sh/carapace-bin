package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var logins_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Gitea logins",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logins_listCmd).Standalone()

	logins_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	loginsCmd.AddCommand(logins_listCmd)

	carapace.Gen(logins_listCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
	})
}
