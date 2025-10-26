package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Print the default `.taplo.toml` configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_defaultCmd).Standalone()

	configCmd.AddCommand(config_defaultCmd)
}
