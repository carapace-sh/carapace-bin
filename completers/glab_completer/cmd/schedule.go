package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:     "schedule <command> [flags]",
	Short:   "Work with GitLab CI schedules",
	Aliases: []string{"sched", "skd"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleCmd).Standalone()

	scheduleCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(scheduleCmd)

	carapace.Gen(scheduleCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(scheduleCmd),
	})
}
