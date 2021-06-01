package cmd

import (
	"github.com/spf13/cobra"
)

var stack_tag_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a stack tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	stack_tagCmd.AddCommand(stack_tag_rmCmd)
}
