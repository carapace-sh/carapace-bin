package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_defaultBranch_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the default branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_defaultBranch_setCmd).Standalone()

	config_state_defaultBranch_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_defaultBranchCmd.AddCommand(config_state_defaultBranch_setCmd)

	carapace.Gen(config_state_defaultBranch_setCmd).PositionalCompletion(
		wt.ActionBranches(),
	)
}
