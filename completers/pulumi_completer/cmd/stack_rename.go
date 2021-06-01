package cmd

import (
	"github.com/spf13/cobra"
)

var stack_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an existing stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	stackCmd.AddCommand(stack_renameCmd)
}
