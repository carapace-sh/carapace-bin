package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all locally installed Node.js versions",
	Aliases: []string{"ls"},
	Run:     func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	rootCmd.AddCommand(listCmd)
}
