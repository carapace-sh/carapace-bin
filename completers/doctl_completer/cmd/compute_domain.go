package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Display commands that manage domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domainCmd).Standalone()
	computeCmd.AddCommand(compute_domainCmd)
}
