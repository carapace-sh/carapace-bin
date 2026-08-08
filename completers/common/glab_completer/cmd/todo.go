package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo <command> [flags]",
	Short: "Manage your to-do list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(todoCmd).Standalone()

	rootCmd.AddCommand(todoCmd)
}
