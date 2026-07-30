package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statisticsCmd = &cobra.Command{
	Use:   "statistics",
	Short: "display server statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statisticsCmd).Standalone()
	rootCmd.AddCommand(statisticsCmd)
}
