package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_recordsCmd = &cobra.Command{
	Use:   "records",
	Short: "Manage DNS records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_recordsCmd).Standalone()
	compute_domainCmd.AddCommand(compute_domain_recordsCmd)
}
