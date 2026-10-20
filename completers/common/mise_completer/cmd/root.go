// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mise",
	Short: "Dev tools, env vars, and tasks in one CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("cd", "C", "", "Run as if mise were started in DIR")
	rootCmd.Flags().BoolP("continue-on-error", "c", false, "Continue running tasks even if one fails")
	rootCmd.PersistentFlags().Bool("debug", false, "Set the log level to debug")
	rootCmd.Flags().BoolP("dry-run", "n", false, "Dry run, don't actually do anything")
	rootCmd.PersistentFlags().StringSliceP("env", "E", nil, "Load the `mise.<ENV>.toml` config files (same as MISE_ENV)")
	rootCmd.Flags().BoolP("force", "f", false, "Force the operation")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().StringP("jobs", "j", "", "How many jobs to run in parallel; defaults to the `jobs` setting")
	rootCmd.PersistentFlags().Bool("locked", false, "Require download URLs from the lockfile when installing")
	rootCmd.PersistentFlags().String("log-level", "", "")
	rootCmd.Flags().Bool("no-config", false, "Do not load any config files")
	rootCmd.Flags().Bool("no-env", false, "Do not load environment variables from config files")
	rootCmd.Flags().Bool("no-hooks", false, "Do not run hooks from config files")
	rootCmd.Flags().Bool("no-timing", false, "Hide the elapsed time printed after each task completes")
	rootCmd.Flags().Bool("no-timings", false, "Hide the elapsed time printed after each task completes")
	rootCmd.Flags().String("output", "", "How task output is displayed when running a task (same as `mise run --output`)")
	rootCmd.PersistentFlags().StringSliceP("profile", "P", nil, "Set the profile (environment)")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress non-error messages")
	rootCmd.PersistentFlags().Bool("raw", false, "Connect tasks and install commands directly to the terminal")
	rootCmd.Flags().StringP("shell", "s", "", "")
	rootCmd.PersistentFlags().Bool("silent", false, "Suppress all task output and mise non-error messages")
	rootCmd.Flags().Bool("timing", false, "Show the elapsed time after each task completes")
	rootCmd.Flags().Bool("timings", false, "Show the elapsed time after each task completes")
	rootCmd.Flags().StringSliceP("tool", "t", nil, "Tools to load in addition to those in mise.toml, such as `node@24 python@3.13`")
	rootCmd.PersistentFlags().Bool("trace", false, "Set the log level to trace")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Show extra output (use -vv for even more)")
	rootCmd.Flags().BoolP("version", "V", false, "")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Answer yes to all confirmation prompts")
	rootCmd.Flag("continue-on-error").Hidden = true
	rootCmd.Flag("debug").Hidden = true
	rootCmd.Flag("dry-run").Hidden = true
	rootCmd.Flag("force").Hidden = true
	rootCmd.Flag("log-level").Hidden = true
	rootCmd.Flag("no-timing").Hidden = true
	rootCmd.Flag("no-timings").Hidden = true
	rootCmd.Flag("output").Hidden = true
	rootCmd.Flag("profile").Hidden = true
	rootCmd.Flag("shell").Hidden = true
	rootCmd.Flag("timing").Hidden = true
	rootCmd.Flag("timings").Hidden = true
	rootCmd.Flag("tool").Hidden = true
	rootCmd.Flag("trace").Hidden = true
	rootCmd.Flag("version").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cd":        carapace.ActionDirectories(),
		"log-level": carapace.ActionValues("trace", "debug", "info", "warn", "error", "warning"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
