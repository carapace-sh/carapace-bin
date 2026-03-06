package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_previousBranch_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the previous branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_previousBranch_setCmd).Standalone()

	config_state_previousBranch_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_previousBranchCmd.AddCommand(config_state_previousBranch_setCmd)

	carapace.Gen(config_state_previousBranch_setCmd).PositionalCompletion(
		wt.ActionBranches(),
	)
}
