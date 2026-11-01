package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cupsfilter",
	Short: "convert files using CUPS filters",
	Long:  "https://keith.github.io/xcode-manpages/cupsfilter.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("D", "D", "", "Delete temp files")
	rootCmd.Flags().StringS("I", "I", "", "Input type")
	rootCmd.Flags().StringS("J", "J", "", "Job ID")
	rootCmd.Flags().StringS("O", "O", "", "Output type")
	rootCmd.Flags().StringS("U", "U", "", "Username")
	rootCmd.Flags().BoolS("h", "h", false, "Print help")
	rootCmd.Flags().StringS("i", "i", "", "Input file")
	rootCmd.Flags().StringS("j", "j", "", "Job title")
	rootCmd.Flags().StringS("m", "m", "", "MIME type")
	rootCmd.Flags().StringS("o", "o", "", "Options")
	rootCmd.Flags().StringS("p", "p", "", "PPD file")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.ActionFiles(),
		"p": carapace.ActionFiles(),
	})
}
