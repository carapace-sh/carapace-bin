package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:     "schedule <command> [flags]",
	Short:   "Work with GitLab CI/CD schedules.",
	Aliases: []string{"sched", "skd"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleCmd).Standalone()

	scheduleCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(scheduleCmd)

	carapace.Gen(scheduleCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(scheduleCmd),
	})
}
