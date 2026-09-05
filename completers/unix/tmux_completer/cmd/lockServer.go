package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lockServerCmd = &cobra.Command{
	Use:     "lock-server",
	Aliases: []string{"lock"},
	Short:   "lock all clients attached to the server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockServerCmd).Standalone()

	rootCmd.AddCommand(lockServerCmd)
}
