package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convert_toTOMLCmd = &cobra.Command{
	Use:   "toTOML",
	Short: "Convert front matter to TOML",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convert_toTOMLCmd).Standalone()
	convertCmd.AddCommand(convert_toTOMLCmd)
}
