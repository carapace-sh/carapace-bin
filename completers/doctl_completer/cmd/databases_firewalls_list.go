package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_firewalls_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of firewall rules for a given database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_firewalls_listCmd).Standalone()
	databases_firewallsCmd.AddCommand(databases_firewalls_listCmd)
}
