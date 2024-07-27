package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeComponentCmd = &cobra.Command{
	Use:   "make:component [-f|--force] [--inline] [--view] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new view component class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeComponentCmd).Standalone()

	makeComponentCmd.Flags().Bool("force", false, "Create the class even if the component already exists")
	makeComponentCmd.Flags().Bool("inline", false, "Create a component that renders an inline view")
	makeComponentCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Component")
	makeComponentCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Component")
	makeComponentCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Component")
	makeComponentCmd.Flags().Bool("view", false, "Create an anonymous component with only a view")
	rootCmd.AddCommand(makeComponentCmd)
}
