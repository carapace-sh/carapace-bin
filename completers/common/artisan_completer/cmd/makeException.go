package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeExceptionCmd = &cobra.Command{
	Use:   "make:exception [-f|--force] [--render] [--report] [--] <name>",
	Short: "Create a new custom exception class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeExceptionCmd).Standalone()

	makeExceptionCmd.Flags().Bool("force", false, "Create the class even if the exception already exists")
	makeExceptionCmd.Flags().Bool("render", false, "Create the exception with an empty render method")
	makeExceptionCmd.Flags().Bool("report", false, "Create the exception with an empty report method")
	rootCmd.AddCommand(makeExceptionCmd)
}
