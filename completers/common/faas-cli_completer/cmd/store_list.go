package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available OpenFaaS functions in a store",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_listCmd).Standalone()
	store_listCmd.Flags().BoolP("json", "j", false, "Output functions as JSON")
	store_listCmd.Flags().BoolP("verbose", "v", true, "Enable verbose output to see the full description of each function in the store")
	storeCmd.AddCommand(store_listCmd)
}
