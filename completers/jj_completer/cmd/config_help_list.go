package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List variables set in config file, along with their values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_listCmd).Standalone()

	config_helpCmd.AddCommand(config_help_listCmd)
}
