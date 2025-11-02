package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "but",
	Short: "A GitButler CLI tool",
	Long:  "https://github.com/gitbutlerapp/gitbutler/tree/master/crates/but",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "branching and committing"},
		&cobra.Group{ID: "editing commits"},
		&cobra.Group{ID: "inspection"},
		&cobra.Group{ID: "operation history"},
	)

	rootCmd.Flags().StringP("current-dir", "C", "", "Run as if gitbutler-cli was started in PATH instead of the current working directory")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("json", "j", false, "Whether to use JSON output format")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"current-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("current-dir")))
	})
}
