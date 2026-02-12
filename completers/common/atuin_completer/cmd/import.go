package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import shell history from file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(importCmd)
}
