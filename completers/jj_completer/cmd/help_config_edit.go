package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_config_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Start an editor on a jj config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_editCmd).Standalone()

	help_configCmd.AddCommand(help_config_editCmd)
}
