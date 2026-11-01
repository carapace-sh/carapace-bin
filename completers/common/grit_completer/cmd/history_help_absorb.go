package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_absorbCmd = &cobra.Command{
	Use:   "absorb",
	Short: "Assign staged hunks into recent commits and replay descendants once",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_absorbCmd).Standalone()

	history_helpCmd.AddCommand(history_help_absorbCmd)
}
