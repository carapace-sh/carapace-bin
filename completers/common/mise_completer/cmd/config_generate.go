package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a mise.toml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_generateCmd).Standalone()

	configCmd.AddCommand(config_generateCmd)
}
