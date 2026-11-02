package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliases_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Output the path to the aliases.json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliases_pathCmd).Standalone()

	aliasesCmd.AddCommand(aliases_pathCmd)
}
