package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listPendingCmd = &cobra.Command{
	Use:     "list-pending",
	Short:   "List pending packages",
	Aliases: []string{"lp"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPendingCmd).Standalone()

	rootCmd.AddCommand(listPendingCmd)
}
