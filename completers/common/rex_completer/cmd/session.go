package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Manage sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionCmd).Standalone()

	rootCmd.AddCommand(sessionCmd)

	sessionCmd.AddCommand(newCmd, lsCmd, attachCmd, killCmd)
}
