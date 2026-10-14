package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domain_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete domain(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domain_deleteCmd).Standalone()

	domainCmd.AddCommand(domain_deleteCmd)
}
