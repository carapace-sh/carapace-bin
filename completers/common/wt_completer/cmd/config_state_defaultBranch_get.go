package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_defaultBranch_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the default branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_defaultBranch_getCmd).Standalone()

	config_state_defaultBranch_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_defaultBranchCmd.AddCommand(config_state_defaultBranch_getCmd)
}
