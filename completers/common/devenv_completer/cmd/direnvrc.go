package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var direnvrcCmd = &cobra.Command{
	Use:   "direnvrc",
	Short: "Print a direnvrc that adds devenv support to direnv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(direnvrcCmd).Standalone()

	rootCmd.AddCommand(direnvrcCmd)
}
