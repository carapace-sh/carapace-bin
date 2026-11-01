package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sfltool",
	Short: "tool for testing and debugging SharedFileList",
	Long:  "https://keith.github.io/xcode-manpages/sfltool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("z", "z", false, "Archive the directory")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("archive"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
