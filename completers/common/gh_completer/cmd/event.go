package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventCmd = &cobra.Command{
	Use:   "event <command>",
	Short: "Listen to GitHub events (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventCmd).Standalone()

	rootCmd.AddCommand(eventCmd)
}
