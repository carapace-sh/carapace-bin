package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var times_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List tracked times on issues & pulls",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(times_listCmd).Standalone()

	times_listCmd.Flags().String("fields", "", "Comma-separated list of fields to print. Available values:")
	times_listCmd.Flags().StringP("from", "f", "", "Show only times tracked after this date")
	times_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	times_listCmd.Flags().BoolP("mine", "m", false, "Show all times tracked by you across all repositories (overrides command arguments)")
	times_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	times_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	times_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	times_listCmd.Flags().BoolP("total", "t", false, "Print the total duration at the end")
	times_listCmd.Flags().StringP("until", "u", "", "Show only times tracked before this date")
	timesCmd.AddCommand(times_listCmd)

	// TODO completion
	carapace.Gen(times_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionTimeFields().UniqueList(","),
		"from":   time.ActionDate(),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"until":  time.ActionDate(),
	})
}
