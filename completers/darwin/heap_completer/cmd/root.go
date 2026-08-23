package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "heap",
	Short: "list all malloc-allocated buffers in a process's heap",
	Long:  "https://keith.github.io/xcode-manpages/heap.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("addresses", false, "Show addresses")
	rootCmd.Flags().Bool("guessNonObjects", false, "Guess non-object allocations")
	rootCmd.Flags().Bool("noContent", false, "No content")
	rootCmd.Flags().BoolS("s", "s", false, "Sort by size")
	rootCmd.Flags().Bool("showSizes", false, "Show sizes")
	rootCmd.Flags().Bool("sortBySize", false, "Sort by size")
	rootCmd.Flags().Bool("sumObjectFields", false, "Sum object fields")
	rootCmd.Flags().BoolS("z", "z", false, "Show zones")
	rootCmd.Flags().Bool("zones", false, "Show zones")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessIds(),
	)
}
