package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fontrestore",
	Short: "restore system fonts to a pristine state",
	Long:  "https://keith.github.io/xcode-manpages/fontrestore.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "Dry run")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("default"),
	)
}
