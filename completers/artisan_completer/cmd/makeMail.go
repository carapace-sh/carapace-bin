package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeMailCmd = &cobra.Command{
	Use:   "make:mail [-f|--force] [-m|--markdown [MARKDOWN]] [--view [VIEW]] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new email class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeMailCmd).Standalone()

	makeMailCmd.Flags().Bool("force", false, "Create the class even if the mailable already exists")
	makeMailCmd.Flags().String("markdown", "", "Create a new Markdown template for the mailable")
	makeMailCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Mailable")
	makeMailCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Mailable")
	makeMailCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Mailable")
	makeMailCmd.Flags().String("view", "", "Create a new Blade template for the mailable")
	rootCmd.AddCommand(makeMailCmd)
}
