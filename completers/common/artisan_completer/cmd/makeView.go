package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeViewCmd = &cobra.Command{
	Use:   "make:view [--extension [EXTENSION]] [-f|--force] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeViewCmd).Standalone()

	makeViewCmd.Flags().String("extension", "", "The extension of the generated view")
	makeViewCmd.Flags().Bool("force", false, "Create the view even if the view already exists")
	makeViewCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the View")
	makeViewCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the View")
	makeViewCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the View")
	rootCmd.AddCommand(makeViewCmd)
}
