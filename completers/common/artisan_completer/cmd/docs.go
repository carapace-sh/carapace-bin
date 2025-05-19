package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs [<page> [<section>]]",
	Short: "Access the Laravel documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsCmd).Standalone()

	rootCmd.AddCommand(docsCmd)
}
