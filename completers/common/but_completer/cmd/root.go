package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "but",
	Short: "A CLI for GitButler",
	Long:  "https://github.com/gitbutlerapp/gitbutler/crates/gitbutler-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("current-dir", "C", "", "Run as if gitbutler-cli was started in PATH instead of the current working directory")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("trace", "d", false, "Enable tracing for debug and performance information printed to stderr")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"current-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("current-dir")))
	})
}
