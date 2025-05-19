package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stack_selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Switch the current workspace to the given stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_selectCmd).Standalone()
	stack_selectCmd.PersistentFlags().BoolP("create", "c", false, "If selected stack does not exist, create it")
	stack_selectCmd.PersistentFlags().String("secrets-provider", "default", "Use with --create flag, The type of the provider that should be used to encrypt and decrypt secrets")
	stackCmd.AddCommand(stack_selectCmd)

	carapace.Gen(stack_selectCmd).PositionalCompletion(
		action.ActionStacks(stack_selectCmd, action.StackOpts{}),
	)
}
