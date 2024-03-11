package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convert_toJSONCmd = &cobra.Command{
	Use:   "toJSON",
	Short: "Convert front matter to JSON",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convert_toJSONCmd).Standalone()
	convertCmd.AddCommand(convert_toJSONCmd)
}
