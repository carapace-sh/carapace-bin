package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_logs_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get log file paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_logs_getCmd).Standalone()

	config_state_logs_getCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_logs_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_logs_getCmd.Flags().String("hook", "", "Get path for a specific log file")
	config_state_logsCmd.AddCommand(config_state_logs_getCmd)

	carapace.Gen(config_state_logs_getCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
		// TODO hook
	})
}
