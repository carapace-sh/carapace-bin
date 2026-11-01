package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "wt",
	Short: "Git worktree management for parallel AI agent workflows",
	Long:  "https://worktrunk.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringS("C", "C", "", "Working directory for this command")
	rootCmd.PersistentFlags().String("config", "", "User config file path")
	rootCmd.PersistentFlags().StringSlice("config-set", nil, "Override config with inline TOML, e.g. --config-set list.full=true (repeatable)")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Verbose output (-v: info logs + hook/alias template variables on stderr; -vv: also debug logs and raw subprocess output written to .git/wt/logs/). Set WORKTRUNK_VERBOSE=0|1|2 to apply the same level everywhere - including shell completion, which no flag can reach")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Skip approval prompts")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C":      carapace.ActionDirectories(),
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("C")))
	})
}
