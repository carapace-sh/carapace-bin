package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update config file to set the given option to a given value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_setCmd).Standalone()

	help_configCmd.AddCommand(help_config_setCmd)
}
