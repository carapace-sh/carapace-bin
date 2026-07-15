package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tab_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tabs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tab_listCmd).Standalone()

	tab_listCmd.Flags().String("workspace", "", "")
	tabCmd.AddCommand(tab_listCmd)
}
