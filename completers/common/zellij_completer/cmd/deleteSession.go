package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteSessionCmd = &cobra.Command{
	Use:     "delete-session",
	Short:   "Delete a specific session",
	Aliases: []string{"d"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteSessionCmd).Standalone()

	deleteSessionCmd.Flags().BoolP("force", "f", false, "Kill the session if it's running before deleting it")
	deleteSessionCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(deleteSessionCmd)
}
