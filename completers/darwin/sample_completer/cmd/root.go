package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sample",
	Short: "profile a process during a time interval",
	Long:  "https://keith.github.io/xcode-manpages/sample.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("e", "e", false, "Executable")
	rootCmd.Flags().StringS("f", "f", "", "File to save output to")
	rootCmd.Flags().BoolS("", "fullPaths", false, "Show full paths")
	rootCmd.Flags().BoolS("", "mayDie", false, "Process may die during sampling")
	rootCmd.Flags().BoolS("", "wait", false, "Wait for the process to start")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessIds(),
	)
}
