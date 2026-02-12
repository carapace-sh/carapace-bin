package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Calculate statistics for your history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_statsCmd).Standalone()

	helpCmd.AddCommand(help_statsCmd)
}
