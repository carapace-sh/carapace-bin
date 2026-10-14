package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var performanceCmd = &cobra.Command{
	Use:   "performance",
	Short: "Apply the performance profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(performanceCmd).Standalone()

	rootCmd.AddCommand(performanceCmd)
}
