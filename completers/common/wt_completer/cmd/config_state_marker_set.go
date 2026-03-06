package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_marker_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set marker for a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_marker_setCmd).Standalone()

	config_state_marker_setCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_marker_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_markerCmd.AddCommand(config_state_marker_setCmd)

	carapace.Gen(config_state_marker_setCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
		// TODO hook
	})
}
