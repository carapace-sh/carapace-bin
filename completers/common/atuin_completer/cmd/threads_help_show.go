package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a thread and its messages. Accepts an id, a harness id or `latest`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_showCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_showCmd)
}
