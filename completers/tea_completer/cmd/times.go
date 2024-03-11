package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var timesCmd = &cobra.Command{
	Use:     "times",
	Short:   "Operate on tracked times of a repository's issues & pulls",
	GroupID: "ENTITIES",
	Aliases: []string{"time", "t"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timesCmd).Standalone()

	timesCmd.Flags().String("fields", "", "Comma-separated list of fields to print. Available values:")
	timesCmd.Flags().StringP("from", "f", "", "Show only times tracked after this date")
	timesCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	timesCmd.Flags().BoolP("mine", "m", false, "Show all times tracked by you across all repositories (overrides command arguments)")
	timesCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	timesCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	timesCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	timesCmd.Flags().BoolP("total", "t", false, "Print the total duration at the end")
	timesCmd.Flags().StringP("until", "u", "", "Show only times tracked before this date")
	rootCmd.AddCommand(timesCmd)

	// TODO completion
	carapace.Gen(timesCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionTimeFields().UniqueList(","),
		"from":   time.ActionDate(),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"until":  time.ActionDate(),
	})
}
