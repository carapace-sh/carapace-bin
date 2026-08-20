package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listBuffersCmd = &cobra.Command{
	Use:     "list-buffers",
	Aliases: []string{"lsb"},
	Short:   "list paste buffers of a session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listBuffersCmd).Standalone()

	listBuffersCmd.Flags().StringS("F", "F", "", "specify output format")
	listBuffersCmd.Flags().StringS("O", "O", "", "initial sort order")
	listBuffersCmd.Flags().StringS("f", "f", "", "filter items")
	listBuffersCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	rootCmd.AddCommand(listBuffersCmd)
}
