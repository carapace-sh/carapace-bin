package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteAllSessionsCmd = &cobra.Command{
	Use:     "delete-all-sessions",
	Short:   "Delete all sessions",
	Aliases: []string{"da"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteAllSessionsCmd).Standalone()

	deleteAllSessionsCmd.Flags().BoolP("force", "f", false, "Kill the sessions if they're running before deleting them")
	deleteAllSessionsCmd.Flags().BoolP("help", "h", false, "Print help")
	deleteAllSessionsCmd.Flags().BoolP("yes", "y", false, "Automatic yes to prompts")
	rootCmd.AddCommand(deleteAllSessionsCmd)
}
