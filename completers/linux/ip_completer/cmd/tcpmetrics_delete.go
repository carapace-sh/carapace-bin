package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tcpmetrics_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete a single TCP metrics entry",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcpmetrics_deleteCmd).Standalone()

	tcpmetricsCmd.AddCommand(tcpmetrics_deleteCmd)
}
