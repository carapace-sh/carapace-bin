package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List users or locked volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().Bool("extended", false, "Show extended information")
	listCmd.Flags().Bool("offline", false, "Display offline, locked volumes")
	listCmd.Flags().Bool("verbose", false, "Enable verbose mode")
}
