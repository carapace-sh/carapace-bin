package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duoCmd = &cobra.Command{
	Use:   "duo <command> prompt",
	Short: "Generate terminal commands from natural language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duoCmd).Standalone()

	rootCmd.AddCommand(duoCmd)
}
