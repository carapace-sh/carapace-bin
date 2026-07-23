package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agent_startCmd = &cobra.Command{
	Use:   "start <name>",
	Short: "Start an agent command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_startCmd).Standalone()

	agent_startCmd.Flags().String("cwd", "", "")
	agent_startCmd.Flags().StringSlice("env", nil, "Set an environment variable for the launched process")
	agent_startCmd.Flags().Bool("focus", false, "")
	agent_startCmd.Flags().Bool("no-focus", false, "")
	agent_startCmd.Flags().String("split", "", "")
	agent_startCmd.Flags().String("tab", "", "")
	agent_startCmd.Flags().String("workspace", "", "")
	agentCmd.AddCommand(agent_startCmd)

	carapace.Gen(agent_startCmd).FlagCompletion(carapace.ActionMap{
		"cwd":   carapace.ActionFiles(),
		"split": carapace.ActionValues("right", "down"),
	})
}
