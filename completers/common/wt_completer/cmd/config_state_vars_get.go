package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_vars_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_vars_getCmd).Standalone()

	config_state_vars_getCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_vars_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_varsCmd.AddCommand(config_state_vars_getCmd)

	carapace.Gen(config_state_vars_getCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
	})

	carapace.Gen(config_state_vars_getCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
