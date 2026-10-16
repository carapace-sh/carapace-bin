package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a mise.toml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_configCmd).Standalone()

	generateCmd.AddCommand(generate_configCmd)
}
