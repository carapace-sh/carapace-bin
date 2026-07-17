package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitNode_connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to most recently used exit node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitNode_connectCmd).Standalone()

	exitNodeCmd.AddCommand(exitNode_connectCmd)
}
