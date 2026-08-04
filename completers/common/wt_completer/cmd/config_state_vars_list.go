package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var config_state_vars_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_vars_listCmd).Standalone()

	config_state_vars_listCmd.Flags().String("branch", "", "Target branch (defaults to current)")
	config_state_vars_listCmd.Flags().String("format", "", "Output format (text, json)")
	config_state_vars_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_varsCmd.AddCommand(config_state_vars_listCmd)

	carapace.Gen(config_state_vars_listCmd).FlagCompletion(carapace.ActionMap{
		"branch": wt.ActionBranches(),
		"format": carapace.ActionValues("text", "json"),
	})
}
