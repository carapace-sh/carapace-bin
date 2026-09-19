package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugnamecompleteCmd = &cobra.Command{
	Use:    "debugnamecomplete",
	Short:  "NAME...",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugnamecompleteCmd).Standalone()

	rootCmd.AddCommand(debugnamecompleteCmd)
}
