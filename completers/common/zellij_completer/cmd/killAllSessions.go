package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killAllSessionsCmd = &cobra.Command{
	Use:     "kill-all-sessions",
	Short:   "Kill all sessions",
	Aliases: []string{"ka"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killAllSessionsCmd).Standalone()

	killAllSessionsCmd.Flags().BoolP("help", "h", false, "Print help")
	killAllSessionsCmd.Flags().BoolP("yes", "y", false, "Automatic yes to prompts")
	rootCmd.AddCommand(killAllSessionsCmd)
}
