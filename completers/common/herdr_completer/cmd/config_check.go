package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate config.toml and print diagnostics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_checkCmd).Standalone()

	configCmd.AddCommand(config_checkCmd)
}
