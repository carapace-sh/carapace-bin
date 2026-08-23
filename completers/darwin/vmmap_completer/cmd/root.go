package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vmmap",
	Short: "Print virtual memory map of a process",
	Long:  "https://keith.github.io/xcode-manpages/vmmap.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allSplitLibs", false, "Show all split library regions")
	rootCmd.Flags().Bool("attributes", false, "Show region attributes")
	rootCmd.Flags().Bool("forkCorpse", false, "Create a corpse for the process")
	rootCmd.Flags().Bool("fullStacks", false, "Show full stack backtraces")
	rootCmd.Flags().Bool("interleaved", false, "Show interleaved output")
	rootCmd.Flags().Bool("noCoalesce", false, "Do not coalesce regions")
	rootCmd.Flags().Bool("pages", false, "Show pages breakdown")
	rootCmd.Flags().BoolP("sortBySize", "s", false, "Sort regions by size")
	rootCmd.Flags().Bool("stacks", false, "Show stack backtraces")
	rootCmd.Flags().Bool("summary", false, "Show summary only")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.Flags().BoolP("wide", "w", false, "Wide output")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues()
		}),
	)
}