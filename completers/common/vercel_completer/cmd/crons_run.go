package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var crons_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Trigger a cron job to run immediately",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crons_runCmd).Standalone()

	crons_runCmd.Flags().String("project", "", "Project name or ID")

	cronsCmd.AddCommand(crons_runCmd)

	carapace.Gen(crons_runCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(crons_runCmd),
	})

	carapace.Gen(crons_runCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
