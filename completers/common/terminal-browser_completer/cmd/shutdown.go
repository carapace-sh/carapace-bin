package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Stop the daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shutdownCmd).Standalone()

	rootCmd.AddCommand(shutdownCmd)
}
