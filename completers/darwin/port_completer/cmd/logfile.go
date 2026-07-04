package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logfileCmd = &cobra.Command{
	Use:   "logfile",
	Short: "Display the path to the log file for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logfileCmd).Standalone()
	rootCmd.AddCommand(logfileCmd)
}
