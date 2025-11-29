package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_lsCmd).Standalone()

	servicesCmd.AddCommand(services_lsCmd)
}
