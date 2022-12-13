package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apm_deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Manage New Relic APM deployment markers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_deploymentCmd).Standalone()
	apmCmd.AddCommand(apm_deploymentCmd)
}
