package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_threadsCmd = &cobra.Command{
	Use:   "threads",
	Short: "Search, read and resume your coding agents' captured threads",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_threadsCmd).Standalone()

	helpCmd.AddCommand(help_threadsCmd)
}
