package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tabCmd = &cobra.Command{
	Use:     "tab",
	Short:   "Edit tab information for installed formulae or casks",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tabCmd).Standalone()

	tabCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	tabCmd.Flags().Bool("debug", false, "Display any debugging information.")
	tabCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	tabCmd.Flags().Bool("help", false, "Show this message.")
	tabCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	tabCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(tabCmd)
}
