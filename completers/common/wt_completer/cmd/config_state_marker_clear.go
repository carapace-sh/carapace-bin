package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_marker_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear marker for a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_marker_clearCmd).Standalone()

	config_state_marker_clearCmd.Flags().Bool("all", false, "Clear all markers")
	config_state_marker_clearCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_marker_clearCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_markerCmd.AddCommand(config_state_marker_clearCmd)

	carapace.Gen(config_state_marker_clearCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
		// TODO hook
	})
}
