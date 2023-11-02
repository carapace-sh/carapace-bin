package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var util_help_configSchemaCmd = &cobra.Command{
	Use:   "config-schema",
	Short: "Print the JSON schema for the jj TOML config format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_help_configSchemaCmd).Standalone()

	util_helpCmd.AddCommand(util_help_configSchemaCmd)
}
