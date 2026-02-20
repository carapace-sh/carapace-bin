package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_editCmd).Standalone()

	config_helpCmd.AddCommand(config_help_editCmd)
}
