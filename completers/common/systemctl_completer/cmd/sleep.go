package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sleepCmd = &cobra.Command{
	Use:     "sleep",
	Short:   "Put the system to sleep",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sleepCmd).Standalone()

	rootCmd.AddCommand(sleepCmd)
}
