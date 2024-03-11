package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/spf13/cobra"
)

var logLevelCmd = &cobra.Command{
	Use:     "log-level",
	Short:   "Get/set logging threshold for manager",
	GroupID: "manager state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logLevelCmd).Standalone()

	rootCmd.AddCommand(logLevelCmd)

	carapace.Gen(logLevelCmd).PositionalCompletion(
		systemctl.ActionLogLevels(),
	)
}
