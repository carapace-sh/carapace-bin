package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_help_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import every thread the agents have on this machine, including those capture leaves alone: the ones nobody has gone on with since it was turned on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_help_importCmd).Standalone()

	ai_session_helpCmd.AddCommand(ai_session_help_importCmd)
}
