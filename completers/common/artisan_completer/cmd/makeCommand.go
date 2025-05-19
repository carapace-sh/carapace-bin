package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeCommandCmd = &cobra.Command{
	Use:   "make:command [-f|--force] [--command [COMMAND]] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new Artisan command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeCommandCmd).Standalone()

	makeCommandCmd.Flags().String("command", "", "The terminal command that will be used to invoke the class")
	makeCommandCmd.Flags().Bool("force", false, "Create the class even if the console command already exists")
	makeCommandCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Console command")
	makeCommandCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Console command")
	makeCommandCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Console command")
	rootCmd.AddCommand(makeCommandCmd)
}
