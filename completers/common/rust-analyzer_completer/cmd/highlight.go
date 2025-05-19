package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var highlightCmd = &cobra.Command{
	Use:   "highlight",
	Short: "Highlight stdin as html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(highlightCmd).Standalone()

	highlightCmd.Flags().Bool("rainbow", false, "Enable rainbow highlighting of identifiers")
	rootCmd.AddCommand(highlightCmd)
}
