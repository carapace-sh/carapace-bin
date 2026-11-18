package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_session_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Turn thread capture on (short for `atuin config enable threads`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_session_enableCmd).Standalone()

	ai_help_sessionCmd.AddCommand(ai_help_session_enableCmd)
}
