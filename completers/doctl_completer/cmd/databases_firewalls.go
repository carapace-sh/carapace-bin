package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_firewallsCmd = &cobra.Command{
	Use:   "firewalls",
	Short: "Display commands to manage firewall rules (called`trusted sources` in the control panel) for database clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_firewallsCmd).Standalone()
	databasesCmd.AddCommand(databases_firewallsCmd)
}
