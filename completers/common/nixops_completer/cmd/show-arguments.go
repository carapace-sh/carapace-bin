package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ShowArgumentsCmd = &cobra.Command{
	Use:   "show-arguments",
	Short: "Show Arguments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ShowArgumentsCmd).Standalone()
	rootCmd.AddCommand(ShowArgumentsCmd)
}
