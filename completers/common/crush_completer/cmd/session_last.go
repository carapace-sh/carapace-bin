package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var session_lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Show most recent session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_lastCmd).Standalone()

	session_lastCmd.Flags().Bool("json", false, "output in JSON format")
	sessionCmd.AddCommand(session_lastCmd)
}
