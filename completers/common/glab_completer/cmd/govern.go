package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var governCmd = &cobra.Command{
	Use:   "govern <command>",
	Short: "Manage AI agent governance. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(governCmd).Standalone()

	rootCmd.AddCommand(governCmd)
}
