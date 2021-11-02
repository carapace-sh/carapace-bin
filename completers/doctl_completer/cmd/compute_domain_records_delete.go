package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_records_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a DNS record",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_records_deleteCmd).Standalone()
	compute_domain_records_deleteCmd.Flags().BoolP("force", "f", false, "Delete record without confirmation prompt")
	compute_domain_recordsCmd.AddCommand(compute_domain_records_deleteCmd)
}
