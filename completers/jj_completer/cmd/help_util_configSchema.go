package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_util_configSchemaCmd = &cobra.Command{
	Use:   "config-schema",
	Short: "Print the JSON schema for the jj TOML config format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_util_configSchemaCmd).Standalone()

	help_utilCmd.AddCommand(help_util_configSchemaCmd)
}
