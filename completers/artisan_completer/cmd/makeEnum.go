package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeEnumCmd = &cobra.Command{
	Use:   "make:enum [-s|--string] [-i|--int] [-f|--force] [--] <name>",
	Short: "Create a new enum",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeEnumCmd).Standalone()

	makeEnumCmd.Flags().Bool("force", false, "Create the enum even if the enum already exists")
	makeEnumCmd.Flags().Bool("int", false, "Generate an integer backed enum.")
	makeEnumCmd.Flags().Bool("string", false, "Generate a string backed enum.")
	rootCmd.AddCommand(makeEnumCmd)
}
