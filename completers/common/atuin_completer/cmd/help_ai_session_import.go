package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_session_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import every thread the agents have on this machine, including those capture leaves alone: the ones nobody has gone on with since it was turned on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_session_importCmd).Standalone()

	help_ai_sessionCmd.AddCommand(help_ai_session_importCmd)
}
