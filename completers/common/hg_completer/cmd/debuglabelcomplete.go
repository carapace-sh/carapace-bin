package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuglabelcompleteCmd = &cobra.Command{
	Use:    "debuglabelcomplete",
	Short:  "LABEL...",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuglabelcompleteCmd).Standalone()

	rootCmd.AddCommand(debuglabelcompleteCmd)
}
