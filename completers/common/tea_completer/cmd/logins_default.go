package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var logins_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Get or Set Default Login",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logins_defaultCmd).Standalone()

	logins_defaultCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	loginsCmd.AddCommand(logins_defaultCmd)

	// TODO completion
	carapace.Gen(logins_defaultCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
	})

	carapace.Gen(logins_defaultCmd).PositionalCompletion(
		tea.ActionLogins(),
	)
}
