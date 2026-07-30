package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "enable/disable tracing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(traceCmd).Standalone()
	rootCmd.AddCommand(traceCmd)
}
