package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qeCmd = &cobra.Command{
	Use:   "qe",
	Short: "query events from an event log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qeCmd).Standalone()
	rootCmd.AddCommand(qeCmd)
}
