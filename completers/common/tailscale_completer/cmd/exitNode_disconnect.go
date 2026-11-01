package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitNode_disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from current exit node, if any",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitNode_disconnectCmd).Standalone()

	exitNodeCmd.AddCommand(exitNode_disconnectCmd)
}
