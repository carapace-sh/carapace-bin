package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured aliases (default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_listCmd).Standalone()

	alias_listCmd.Flags().BoolP("help", "h", false, "Print help")
	aliasCmd.AddCommand(alias_listCmd)
}
