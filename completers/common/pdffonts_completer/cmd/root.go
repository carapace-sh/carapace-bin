package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdffonts",
	Short: "Portable Document Format (PDF) font analyzer (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdffonts.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().StringS("f", "f", "", "first page to examine")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().StringS("l", "l", "", "last page to examine")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("subst", "subst", false, "show font substitutions")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
