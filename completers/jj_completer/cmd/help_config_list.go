package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_config_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List variables set in config file, along with their values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_listCmd).Standalone()

	help_configCmd.AddCommand(help_config_listCmd)
}
