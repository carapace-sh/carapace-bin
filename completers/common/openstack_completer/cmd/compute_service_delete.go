package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compute_service_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete compute service(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_service_deleteCmd).Standalone()

	compute_serviceCmd.AddCommand(compute_service_deleteCmd)
}
