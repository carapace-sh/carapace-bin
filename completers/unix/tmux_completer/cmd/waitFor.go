package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waitForCmd = &cobra.Command{
	Use:     "wait-for",
	Aliases: []string{"wait"},
	Short:   "wait for an event or trigger it",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForCmd).Standalone()

	waitForCmd.Flags().BoolS("E", "E", false, "wait for the next event with name")
	waitForCmd.Flags().StringS("F", "F", "", "format must also be true")
	waitForCmd.Flags().BoolS("L", "L", false, "lock the channel")
	waitForCmd.Flags().BoolS("S", "S", false, "signal the channel")
	waitForCmd.Flags().BoolS("U", "U", false, "unlock the channel")
	waitForCmd.Flags().BoolS("l", "l", false, "list the waiters for name")
	waitForCmd.Flags().BoolS("v", "v", false, "print event payload keys")
	waitForCmd.Flags().StringS("w", "w", "", "wake waiter on name immediately")
	rootCmd.AddCommand(waitForCmd)

	carapace.Gen(waitForCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
