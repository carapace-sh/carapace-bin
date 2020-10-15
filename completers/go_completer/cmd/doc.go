package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "show documentation for package or symbol",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().Bool("all", false, "Show all the documentation for the package")
	docCmd.Flags().Bool("cmd", false, "Treat a command like a regular package")
	docCmd.Flags().Bool("short", false, "One-line representation for each symbol")
	docCmd.Flags().Bool("src", false, "Show the full source code for the symbol")
	rootCmd.AddCommand(docCmd)
}
