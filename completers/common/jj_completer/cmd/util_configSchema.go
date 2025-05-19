package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_configSchemaCmd = &cobra.Command{
	Use:   "config-schema",
	Short: "Print the JSON schema for the jj TOML config format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_configSchemaCmd).Standalone()

	util_configSchemaCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_configSchemaCmd)
}
