package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage the repository in a service process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_serviceCmd).Standalone()

	helpCmd.AddCommand(help_serviceCmd)
}
