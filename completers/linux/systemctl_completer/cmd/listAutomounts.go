package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listAutomountsCmd = &cobra.Command{
	Use:     "list-automounts",
	Short:   "List automount units currently in memory, ordered by path",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listAutomountsCmd).Standalone()

	rootCmd.AddCommand(listAutomountsCmd)
}
