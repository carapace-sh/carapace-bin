package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete service(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_deleteCmd).Standalone()

	serviceCmd.AddCommand(service_deleteCmd)
}
