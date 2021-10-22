package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a domain from your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_deleteCmd).Standalone()
	compute_domain_deleteCmd.Flags().BoolP("force", "f", false, "Delete domain without confirmation prompt")
	compute_domainCmd.AddCommand(compute_domain_deleteCmd)
}
