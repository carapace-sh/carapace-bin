package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeScopeCmd = &cobra.Command{
	Use:   "make:scope [-f|--force] [--] <name>",
	Short: "Create a new scope class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeScopeCmd).Standalone()

	makeScopeCmd.Flags().Bool("force", false, "Create the class even if the scope already exists")
	rootCmd.AddCommand(makeScopeCmd)
}
