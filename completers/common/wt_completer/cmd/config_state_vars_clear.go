package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_vars_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear a key or all keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_vars_clearCmd).Standalone()

	config_state_vars_clearCmd.Flags().Bool("all", false, "Clear all keys for the branch")
	config_state_vars_clearCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_vars_clearCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_varsCmd.AddCommand(config_state_vars_clearCmd)

	carapace.Gen(config_state_vars_clearCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
	})
}
