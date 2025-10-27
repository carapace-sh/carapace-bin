package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_switchCmd = &cobra.Command{
	Use:   "switch <stack-name>",
	Short: "Switch between stacks. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_switchCmd).Standalone()

	stackCmd.AddCommand(stack_switchCmd)

	// TODO complete stack names
}
