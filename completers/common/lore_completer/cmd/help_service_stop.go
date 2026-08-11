package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_service_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop service for a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_service_stopCmd).Standalone()

	help_serviceCmd.AddCommand(help_service_stopCmd)
}
