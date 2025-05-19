package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convert_toYAMLCmd = &cobra.Command{
	Use:   "toYAML",
	Short: "Convert front matter to YAML",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convert_toYAMLCmd).Standalone()
	convertCmd.AddCommand(convert_toYAMLCmd)
}
