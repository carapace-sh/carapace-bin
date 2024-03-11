package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask <command> prompt",
	Short: "Generate terminal commands from natural language. (Experimental.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(askCmd).Standalone()

	rootCmd.AddCommand(askCmd)
}
