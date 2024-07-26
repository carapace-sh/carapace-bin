package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaStubsCmd = &cobra.Command{
	Use:   "nova:stubs [--force]",
	Short: "Publish all stubs that are available for customization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaStubsCmd).Standalone()

	novaStubsCmd.Flags().Bool("force", false, "Overwrite any existing files")
	rootCmd.AddCommand(novaStubsCmd)
}
