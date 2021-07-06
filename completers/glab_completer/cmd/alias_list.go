package cmd

import (
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available aliases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	aliasCmd.AddCommand(alias_listCmd)
}
