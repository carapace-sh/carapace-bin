package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hybridSleepCmd = &cobra.Command{
	Use:     "hybrid-sleep",
	Short:   "Hibernate and suspend the system",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hybridSleepCmd).Standalone()

	rootCmd.AddCommand(hybridSleepCmd)
}
