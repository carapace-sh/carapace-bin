package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tomlTestCmd = &cobra.Command{
	Use:   "toml-test",
	Short: "Start a decoder for `toml-test` (https://github.com/BurntSushi/toml-test)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tomlTestCmd).Standalone()

	rootCmd.AddCommand(tomlTestCmd)
}
