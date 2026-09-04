package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_listCmd).Standalone()
	alias_listCmd.Flags().BoolP("verbose", "v", false, "show alias descriptions")
	aliasCmd.AddCommand(alias_listCmd)
}
