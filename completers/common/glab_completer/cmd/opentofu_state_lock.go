package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opentofu_state_lockCmd = &cobra.Command{
	Use:   "lock <state>",
	Short: "Lock the given state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_state_lockCmd).Standalone()

	opentofu_stateCmd.AddCommand(opentofu_state_lockCmd)
}
