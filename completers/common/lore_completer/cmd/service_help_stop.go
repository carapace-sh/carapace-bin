package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_help_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop service for a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_help_stopCmd).Standalone()

	service_helpCmd.AddCommand(service_help_stopCmd)
}
