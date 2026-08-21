package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listPendingCmd = &cobra.Command{
	Use:     "list-pending",
	Aliases: []string{"lp"},
	Short:   "show all packages currently in a state of required configuration",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPendingCmd).Standalone()

	rootCmd.AddCommand(listPendingCmd)
}
