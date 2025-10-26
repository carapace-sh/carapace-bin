package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opentofu_state_unlockCmd = &cobra.Command{
	Use:   "unlock <state>",
	Short: "Unlock the given state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_state_unlockCmd).Standalone()

	opentofu_stateCmd.AddCommand(opentofu_state_unlockCmd)
}
