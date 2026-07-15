package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var session_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_listCmd).Standalone()

	session_listCmd.Flags().Bool("json", false, "")
	sessionCmd.AddCommand(session_listCmd)
}
