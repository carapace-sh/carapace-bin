package cmd

import (
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	aliasCmd.AddCommand(alias_listCmd)
}
