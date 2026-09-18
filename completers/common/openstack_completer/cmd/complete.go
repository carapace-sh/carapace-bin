package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "print bash completion command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completeCmd).Standalone()

	completeCmd.Flags().String("name", "", "Command name to support with command completion")
	completeCmd.Flags().String("shell", "", "Shell being used.")
	rootCmd.AddCommand(completeCmd)
}
