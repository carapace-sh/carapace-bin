package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killServerCmd = &cobra.Command{
	Use:   "kill-server",
	Short: "kill clients, sessions and server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killServerCmd).Standalone()

	rootCmd.AddCommand(killServerCmd)
}
