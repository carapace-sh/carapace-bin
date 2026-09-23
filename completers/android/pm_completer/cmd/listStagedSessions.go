package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listStagedSessionsCmd = &cobra.Command{
	Use:   "staged-sessions",
	Short: "Print all staged sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listStagedSessionsCmd).Standalone()

	listStagedSessionsCmd.Flags().Bool("only-parent", false, "hide all children sessions")
	listStagedSessionsCmd.Flags().Bool("only-ready", false, "show only staged sessions that are ready")
	listStagedSessionsCmd.Flags().Bool("only-sessionid", false, "show only sessionId of each session")

	listCmd.AddCommand(listStagedSessionsCmd)

}
