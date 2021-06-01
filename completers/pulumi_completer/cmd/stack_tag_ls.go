package cmd

import (
	"github.com/spf13/cobra"
)

var stack_tag_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all stack tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	stack_tag_lsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	stack_tagCmd.AddCommand(stack_tag_lsCmd)
}
