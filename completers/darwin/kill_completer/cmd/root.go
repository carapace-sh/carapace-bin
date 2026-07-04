package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kill",
	Short: "terminate or signal a process",
	Long:  "https://keith.github.io/xcode-manpages/kill.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("l", "l", false, "List signal names")

	rootCmd.Flags().StringS("s", "s", "", "Send the specified signal")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"s": ps.ActionKillSignals(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(ps.ActionProcessIds())
}
