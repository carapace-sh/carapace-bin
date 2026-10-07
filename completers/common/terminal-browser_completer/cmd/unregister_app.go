package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unregisterAppCmd = &cobra.Command{
	Use:   "unregister-app",
	Short: "Unregister a terminal-browser application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unregisterAppCmd).Standalone()

	rootCmd.AddCommand(unregisterAppCmd)
}
