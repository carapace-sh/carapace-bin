package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_session_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Turn thread capture on (short for `atuin config enable threads`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_session_enableCmd).Standalone()

	help_ai_sessionCmd.AddCommand(help_ai_session_enableCmd)
}
