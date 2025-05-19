package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	config_setCmd.Flags().StringP("key", "k", "", "the key to set")
	config_setCmd.Flags().StringP("value", "v", "", "the value to set")
	configCmd.AddCommand(config_setCmd)
}
