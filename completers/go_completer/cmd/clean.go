package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove object files and cached files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("cache", false, "remove the entire go build cache")
	cleanCmd.Flags().BoolS("i", "i", false, "remove the corresponding installed archive or binary")
	cleanCmd.Flags().Bool("modcache", false, "remove the entire module download cache")
	cleanCmd.Flags().BoolS("n", "n", false, "only print the remove commands that would be executed")
	cleanCmd.Flags().BoolS("r", "r", false, "apply recursively to all the dependencies")
	cleanCmd.Flags().Bool("testcache", false, "expire all test results in the go build cache")
	cleanCmd.Flags().BoolS("x", "x", false, "print remove commands as being executed")
	rootCmd.AddCommand(cleanCmd)
}
