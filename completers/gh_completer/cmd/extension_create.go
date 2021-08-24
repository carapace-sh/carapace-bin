package cmd

import (
	"github.com/spf13/cobra"
)

var extension_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	extensionCmd.AddCommand(extension_createCmd)
}
