package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_service_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service for a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_service_startCmd).Standalone()

	help_serviceCmd.AddCommand(help_service_startCmd)
}
