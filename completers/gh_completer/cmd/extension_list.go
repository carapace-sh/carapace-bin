package cmd

import (
	"github.com/spf13/cobra"
)

var extension_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed extension commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	extensionCmd.AddCommand(extension_listCmd)
}
