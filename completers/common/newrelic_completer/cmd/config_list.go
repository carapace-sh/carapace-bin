package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the current configuration values",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()
	configCmd.AddCommand(config_listCmd)
}
