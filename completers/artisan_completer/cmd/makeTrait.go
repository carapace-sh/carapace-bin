package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeTraitCmd = &cobra.Command{
	Use:   "make:trait [-f|--force] [--] <name>",
	Short: "Create a new trait",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeTraitCmd).Standalone()

	makeTraitCmd.Flags().Bool("force", false, "Create the trait even if the trait already exists")
	rootCmd.AddCommand(makeTraitCmd)
}
