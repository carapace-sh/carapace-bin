package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugcompleteCmd = &cobra.Command{
	Use:    "debugcomplete",
	Short:  "returns the completion list associated with the given command",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugcompleteCmd).Standalone()

	debugcompleteCmd.Flags().BoolP("options", "o", false, "show the command options")
	rootCmd.AddCommand(debugcompleteCmd)
}
