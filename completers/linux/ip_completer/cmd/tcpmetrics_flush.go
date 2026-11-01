package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tcpmetrics_flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "flush TCP metrics entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcpmetrics_flushCmd).Standalone()

	tcpmetricsCmd.AddCommand(tcpmetrics_flushCmd)
}
