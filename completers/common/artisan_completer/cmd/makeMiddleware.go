package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeMiddlewareCmd = &cobra.Command{
	Use:   "make:middleware [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new middleware class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeMiddlewareCmd).Standalone()

	makeMiddlewareCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Middleware")
	makeMiddlewareCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Middleware")
	makeMiddlewareCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Middleware")
	rootCmd.AddCommand(makeMiddlewareCmd)
}
