package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugrequiresCmd = &cobra.Command{
	Use:    "debugrequires",
	Short:  "print the current repo requirements",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrequiresCmd).Standalone()

	rootCmd.AddCommand(debugrequiresCmd)
}
