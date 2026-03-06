package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_marker_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get marker for a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_marker_getCmd).Standalone()

	config_state_marker_getCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_marker_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_markerCmd.AddCommand(config_state_marker_getCmd)

	carapace.Gen(config_state_marker_getCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
		// TODO hook
	})
}
