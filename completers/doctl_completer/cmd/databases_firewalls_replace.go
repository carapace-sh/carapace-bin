package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_firewalls_replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replaces the firewall rules for a given database. The rules passed in to the --rules flag will replace the firewall rules previously assigned to the database,",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_firewalls_replaceCmd).Standalone()
	databases_firewalls_replaceCmd.Flags().StringSlice("rule", []string{}, "A comma-separated list of firewall rules of format type:value, e.g.: `type:value` (required)")
	databases_firewallsCmd.AddCommand(databases_firewalls_replaceCmd)
}
