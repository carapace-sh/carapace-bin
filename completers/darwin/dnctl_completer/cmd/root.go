package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnctl",
	Short: "traffic shaper control program",
	Long:  "https://keith.github.io/xcode-manpages/dnctl.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Display all entries")
	rootCmd.Flags().BoolS("f", "f", false, "Flush all queues and pipes")
	rootCmd.Flags().BoolS("n", "n", false, "Show numeric addresses only")
	rootCmd.Flags().StringS("p", "p", "", "Preprocessor to use")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet")
	rootCmd.Flags().BoolS("s", "s", false, "Show statistics")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("list", "show", "flush", "delete", "pipe", "queue"),
	)
}
