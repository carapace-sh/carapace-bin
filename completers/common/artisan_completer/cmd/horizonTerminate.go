package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonTerminateCmd = &cobra.Command{
	Use:   "horizon:terminate [--wait]",
	Short: "Terminate the master supervisor so it can be restarted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonTerminateCmd).Standalone()

	horizonTerminateCmd.Flags().Bool("wait", false, "Wait for all workers to terminate")
	rootCmd.AddCommand(horizonTerminateCmd)
}
