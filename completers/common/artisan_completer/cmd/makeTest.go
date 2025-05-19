package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeTestCmd = &cobra.Command{
	Use:   "make:test [-f|--force] [-u|--unit] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new test class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeTestCmd).Standalone()

	makeTestCmd.Flags().Bool("force", false, "Create the test even if the test already exists")
	makeTestCmd.Flags().Bool("pest", false, "Create a Pest test")
	makeTestCmd.Flags().Bool("phpunit", false, "Create a PHPUnit test")
	makeTestCmd.Flags().Bool("unit", false, "Create a unit test")
	rootCmd.AddCommand(makeTestCmd)
}
