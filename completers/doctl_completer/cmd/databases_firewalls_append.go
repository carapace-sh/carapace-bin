package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_firewalls_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Add a database firewall rule to a given database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_firewalls_appendCmd).Standalone()
	databases_firewalls_appendCmd.Flags().String("rule", "", " (required)")
	databases_firewallsCmd.AddCommand(databases_firewalls_appendCmd)
}
