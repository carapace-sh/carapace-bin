package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Parse and show log files for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()
	logCmd.Flags().String("phase", "", "Filter log by phase")
	logCmd.Flags().String("verbosity", "", "Filter log by verbosity level")
	rootCmd.AddCommand(logCmd)
}
