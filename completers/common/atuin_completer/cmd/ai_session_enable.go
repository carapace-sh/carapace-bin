package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Turn thread capture on (short for `atuin config enable threads`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_enableCmd).Standalone()

	ai_session_enableCmd.Flags().BoolP("help", "h", false, "Print help")
	ai_sessionCmd.AddCommand(ai_session_enableCmd)
}
