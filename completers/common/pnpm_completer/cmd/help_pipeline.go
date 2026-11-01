package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Runs a named pipeline of workspace tasks the way a CI run would: a frozen install, affected-since-base selection, the task graph in dependency order without bailing, and cached task results restored instead of re-run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pipelineCmd).Standalone()

	helpCmd.AddCommand(help_pipelineCmd)
}
