package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeProviderCmd = &cobra.Command{
	Use:   "make:provider [-f|--force] [--] <name>",
	Short: "Create a new service provider class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeProviderCmd).Standalone()

	makeProviderCmd.Flags().Bool("force", false, "Create the class even if the provider already exists")
	rootCmd.AddCommand(makeProviderCmd)
}
