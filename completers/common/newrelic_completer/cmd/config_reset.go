package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_resetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Reset a configuration value to its default",
	Aliases: []string{"rm", "delete"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_resetCmd).Standalone()
	config_resetCmd.Flags().StringP("key", "k", "", "the key to delete")
	configCmd.AddCommand(config_resetCmd)
}
