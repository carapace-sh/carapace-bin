package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of a given config option.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_getCmd).Standalone()

	config_helpCmd.AddCommand(config_help_getCmd)
}
