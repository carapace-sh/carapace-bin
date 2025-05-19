package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "tig",
	Short: "text-mode interface for Git",
	Long:  "https://jonas.github.io/tig/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "Start in <path>")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and exit")
	rootCmd.Flags().BoolP("version", "v", false, "Show version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			git.ActionRefs(git.RefOption{}.Default()).UnlessF(condition.CompletingPath),
		).ToA(),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		if f := rootCmd.Flag("C"); f != flag && f.Changed {
			return action.Chdir(f.Value.String())
		}
		return action
	})
}
