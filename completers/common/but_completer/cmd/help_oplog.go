package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_oplogCmd = &cobra.Command{
	Use:   "oplog",
	Short: "Show operation history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_oplogCmd).Standalone()

	helpCmd.AddCommand(help_oplogCmd)
}
