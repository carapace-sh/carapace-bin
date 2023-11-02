package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_help_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Start an editor on a jj config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_editCmd).Standalone()

	config_helpCmd.AddCommand(config_help_editCmd)
}
