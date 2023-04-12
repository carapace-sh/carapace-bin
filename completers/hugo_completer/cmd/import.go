package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import your site from others.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()
	rootCmd.AddCommand(importCmd)
}
