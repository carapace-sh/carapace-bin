package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_stack_setCmd = &cobra.Command{
	Use:   "set REPO STACK",
	Short: "Set Git stack for repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_stack_setCmd).Standalone()

	git_stackCmd.AddCommand(git_stack_setCmd)
}
