package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate files for mise",
	Aliases: []string{"gen"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	rootCmd.AddCommand(generateCmd)
}
