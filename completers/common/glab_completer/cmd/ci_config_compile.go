package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_config_compileCmd = &cobra.Command{
	Use:   "compile [<path>]",
	Short: "View the merged CI/CD configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_config_compileCmd).Standalone()

	ci_configCmd.AddCommand(ci_config_compileCmd)
}
