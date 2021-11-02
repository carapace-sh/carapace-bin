package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_firewalls_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a firewall rule for a given database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_firewalls_removeCmd).Standalone()
	databases_firewalls_removeCmd.Flags().String("uuid", "", " (required)")
	databases_firewallsCmd.AddCommand(databases_firewalls_removeCmd)
}
