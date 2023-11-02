package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of a given config option.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_getCmd).Standalone()

	help_configCmd.AddCommand(help_config_getCmd)
}
