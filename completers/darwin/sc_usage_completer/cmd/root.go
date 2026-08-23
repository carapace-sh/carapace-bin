package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sc_usage",
	Short: "show system call usage statistics",
	Long:  "https://keith.github.io/xcode-manpages/sc_usage.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("E", "E", "", "Execute command")
	rootCmd.Flags().StringS("c", "c", "", "Codefile path")
	rootCmd.Flags().BoolS("e", "e", false, "Sort by call count")
	rootCmd.Flags().BoolS("l", "l", false, "Scrolling output")
	rootCmd.Flags().StringS("s", "s", "", "Sampling interval")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.Batch(
		ps.ActionProcessIds(),
		carapace.ActionFiles(),
	).ToA())
}
