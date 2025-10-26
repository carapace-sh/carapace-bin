package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Print the JSON schema of the `.taplo.toml` configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_schemaCmd).Standalone()

	configCmd.AddCommand(config_schemaCmd)
}
