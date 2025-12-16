package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "terramate",
	Short: "A tool for managing terraform stacks",
	Long:  "https://github.com/mineiros-io/terramate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("changed", "c", false, "Filter by changed infrastructure")
	rootCmd.Flags().StringP("chdir", "C", "", "Sets working directory")
	rootCmd.Flags().Bool("disable-check-git-uncommitted", false, "Disable git check for uncommitted files")
	rootCmd.Flags().Bool("disable-check-git-untracked", false, "Disable git check for untracked files")
	rootCmd.Flags().Bool("disable-checkpoint", false, "Disable checkpoint checks for updates")
	rootCmd.Flags().Bool("disable-checkpoint-signature", false, "Disable checkpoint signature")
	rootCmd.Flags().StringP("git-change-base", "B", "", "Git base ref for computing changes")
	rootCmd.Flags().BoolP("help", "h", false, "Show context-sensitive help.")
	rootCmd.Flags().String("log-destination", "", "Destination of log messages")
	rootCmd.Flags().String("log-fmt", "", "Log format to use: 'console', 'text', or 'json'")
	rootCmd.Flags().String("log-level", "", "Log level to use: 'disabled', 'trace', 'debug', 'info', 'warn', 'error', or 'fatal'")
	rootCmd.Flags().StringSlice("no-tags", nil, "Filter stacks that do not have the given tags")
	rootCmd.Flags().Bool("quiet", false, "Disable output")
	rootCmd.Flags().StringSlice("tags", nil, "Filter stacks by tags. Use \":\" for logical AND and \",\" for logical OR. Example: --tags app:prod filters stacks containing tag \"app\" AND \"prod\". If multiple --tags are provided, an OR expression is created. Example: \"--tags a --tags b\" is the same as \"--tags a,b\"")
	rootCmd.Flags().StringSliceP("verbose", "v", nil, "Increase verboseness of output")
	rootCmd.Flags().Bool("version", false, "Terramate version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chdir":           carapace.ActionDirectories(),
		"git-change-base": git.ActionRefs(git.RefOption{}.Default()),
		"log-destination": carapace.ActionValues("stderr", "stdout"),
		"log-fmt":         carapace.ActionValues("console", "text", "json"),
		"log-level":       carapace.ActionValues("disabled", "trace", "debug", "info", "warn", "error", "fatal").StyleF(style.ForLogLevel),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("chdir").Value.String())
	})
}
