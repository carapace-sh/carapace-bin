package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_addCmd = &cobra.Command{
	Use:   "add",
	Short: "create a new alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_addCmd).Standalone()
	aliasCmd.AddCommand(alias_addCmd)
}
