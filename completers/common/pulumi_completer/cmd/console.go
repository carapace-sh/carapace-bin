package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Opens the current stack in the Pulumi Console",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()
	consoleCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to view")
	rootCmd.AddCommand(consoleCmd)

	carapace.Gen(consoleCmd).FlagCompletion(carapace.ActionMap{
		"stack": action.ActionStacks(consoleCmd, action.StackOpts{}),
	})
}
