package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaLensCmd = &cobra.Command{
	Use:   "nova:lens <name>",
	Short: "Create a new lens class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaLensCmd).Standalone()

	rootCmd.AddCommand(novaLensCmd)
}
