package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "run a command in the named network namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_execCmd).Standalone()

	netnsCmd.AddCommand(netns_execCmd)

	carapace.Gen(netns_execCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
