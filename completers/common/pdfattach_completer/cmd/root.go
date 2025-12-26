package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfattach",
	Short: "Portable Document Format (PDF) document embedded file creator (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdfattach.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("replace", "replace", false, "replace embedded file with same name (if it exists)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
