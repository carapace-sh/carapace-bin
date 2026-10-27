package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_receivePackCmd = &cobra.Command{
	Use:    "receive-pack",
	Short:  "Accept a push over stdin/stdout (run by ssh and grit-http-server)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_receivePackCmd).Standalone()

	helpCmd.AddCommand(help_receivePackCmd)
}
