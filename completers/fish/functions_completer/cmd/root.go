package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "functions",
	Short: "Print or erase functions",
	Long:  "https://fishshell.com/docs/current/cmds/functions.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "list all functions")
	rootCmd.Flags().Bool("all", false, "list all functions")
	rootCmd.Flags().StringS("c", "c", "", "copy a function")
	rootCmd.Flags().String("copy", "", "copy a function")
	rootCmd.Flags().String("color", "", "color output")
	rootCmd.Flags().BoolS("D", "D", false, "report function definition path")
	rootCmd.Flags().Bool("details", false, "report function definition path")
	rootCmd.Flags().StringS("d", "d", "", "change function description")
	rootCmd.Flags().String("description", "", "change function description")
	rootCmd.Flags().BoolS("e", "e", false, "erase functions")
	rootCmd.Flags().Bool("erase", false, "erase functions")
	rootCmd.Flags().BoolS("H", "H", false, "show event handlers")
	rootCmd.Flags().Bool("handlers", false, "show event handlers")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("n", "n", false, "list function names")
	rootCmd.Flags().Bool("names", false, "list function names")
	rootCmd.Flags().Bool("no-details", false, "turn off path reporting")
	rootCmd.Flags().BoolS("q", "q", false, "test if functions exist")
	rootCmd.Flags().Bool("query", false, "test if functions exist")
	rootCmd.Flags().StringS("t", "t", "", "filter handlers by type")
	rootCmd.Flags().String("handlers-type", "", "filter handlers by type")
	rootCmd.Flags().BoolS("v", "v", false, "verbose output")
	rootCmd.Flags().Bool("verbose", false, "verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "never", "auto"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionFunctions(true).FilterArgs(),
	)
}
