package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "pnpm",
	Short: "Fast, disk space efficient package manager",
	Long:  "https://pnpm.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "manage", Title: "Manage Commands"},
		&cobra.Group{ID: "review", Title: "Review Commands"},
		&cobra.Group{ID: "run", Title: "Run Commands"},
		&cobra.Group{ID: "store", Title: "Store Commands"},
		&cobra.Group{ID: "other", Title: "Other Commands"},
	)

	rootCmd.PersistentFlags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	rootCmd.PersistentFlags().Bool("color", false, "Controls colors in the output")
	rootCmd.PersistentFlags().StringP("dir", "C", "", "Change to directory <dir>")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Output usage information")
	rootCmd.PersistentFlags().String("loglevel", "", "What level of logs to report")
	rootCmd.PersistentFlags().Bool("no-color", false, "Controls colors in the output")
	rootCmd.PersistentFlags().Bool("stream", false, "Stream output from child processes immediately")
	rootCmd.PersistentFlags().Bool("use-stderr", false, "Divert all output to stderr")
	rootCmd.PersistentFlags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.Flags().BoolP("version", "v", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":      carapace.ActionDirectories(),
		"loglevel": pnpm.ActionLoglevels(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("dir")))
	})
}
