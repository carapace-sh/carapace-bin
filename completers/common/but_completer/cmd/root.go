package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/util"
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
		&cobra.Group{ID: "inspection"},
		&cobra.Group{ID: "branching and committing"},
		&cobra.Group{ID: "server interactions"},
		&cobra.Group{ID: "editing commits"},
		&cobra.Group{ID: "operation history"},
	)

	rootCmd.Flags().StringP("current-dir", "C", "", "Run as if gitbutler-cli was started in PATH instead of the current working directory")
	rootCmd.Flags().StringP("format", "f", "", "Explicitly control how output should be formatted")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Whether to use JSON output format")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"current-dir": carapace.ActionDirectories(),
		"format":      carapace.ActionValues("human", "shell", "json", "none"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{}.Default()),
			carapace.ActionDirectories(),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Args[0]) {
				return carapace.ActionValues()
			}
			return but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true})
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("current-dir")))
	})
}
