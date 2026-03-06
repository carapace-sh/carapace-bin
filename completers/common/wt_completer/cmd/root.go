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
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Verbose output (-v: hooks, templates; -vv: debug report)")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C":      carapace.ActionDirectories(),
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("C")))
	})
}
