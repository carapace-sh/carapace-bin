package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clearDebugAppCmd = &cobra.Command{
	Use:   "clear-debug-app",
	Short: "Clear the previously set-debug-app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearDebugAppCmd).Standalone()

	rootCmd.AddCommand(clearDebugAppCmd)

}
