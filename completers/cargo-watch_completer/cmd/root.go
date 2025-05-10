package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/cargo"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-watch",
	Short: "Watches over your Cargo project’s source",
	Long:  "https://github.com/watchexec/cargo-watch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("B", "B", "", "Inject RUST_BACKTRACE=VALUE (generally you want to set it to 1) into the environment")
	rootCmd.Flags().StringS("L", "L", "", "Inject RUST_LOG=VALUE into the environment --use-shell <use-shell>")
	rootCmd.Flags().BoolP("clear", "c", false, "Clear the screen before each run")
	rootCmd.Flags().Bool("debug", false, "Show debug output")
	rootCmd.Flags().StringP("delay", "d", "", "File updates debounce delay in seconds [default: 0.5]")
	rootCmd.Flags().StringArrayP("env", "E", nil, "Set environment variables for the command")
	rootCmd.Flags().StringArray("env-file", nil, "Set environment variables from a .env file")
	rootCmd.Flags().StringArrayP("exec", "x", nil, "Cargo command(s) to execute on changes [default: check]")
	rootCmd.Flags().String("features", "", "List of features passed to cargo invocations")
	rootCmd.Flags().BoolP("help", "h", false, "Display this message")
	rootCmd.Flags().StringArrayP("ignore", "i", nil, "Ignore a glob/gitignore-style pattern")
	rootCmd.Flags().Bool("ignore-nothing", false, "Ignore nothing, not even target/ and .git/")
	rootCmd.Flags().Bool("no-dot-ignores", false, "Don’t use .ignore files")
	rootCmd.Flags().Bool("no-process-group", false, "Do not use a process group when running the command")
	rootCmd.Flags().Bool("no-restart", false, "Don’t restart command while it’s still running")
	rootCmd.Flags().Bool("no-vcs-ignores", false, "Don’t use .gitignore files")
	rootCmd.Flags().BoolP("notify", "N", false, "Send a desktop notification when watchexec notices a change (experimental, behaviour may change)")
	rootCmd.Flags().Bool("poll", false, "Force use of polling for file changes")
	rootCmd.Flags().Bool("postpone", false, "Postpone first run until a file changes")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress output from cargo-watch itself")
	rootCmd.Flags().StringArrayP("shell", "s", nil, "Shell command(s) to execute on changes")
	rootCmd.Flags().Bool("skip-local-deps", false, "Don't try to find local dependencies of the current crate and watch their working directories. Only watch the current directory.")
	rootCmd.Flags().String("use-shell", "", "Use a different shell. E.g. --use-shell=bash")
	rootCmd.Flags().BoolP("version", "V", false, "Display version information")
	rootCmd.Flags().StringArrayP("watch", "w", nil, "Watch specific file(s) or folder(s). Disables finding and watching local dependencies.")
	rootCmd.Flags().Bool("watch-when-idle", false, "Ignore events emitted while the commands run.")
	rootCmd.Flags().Bool("why", false, "Show paths that changed")
	rootCmd.Flags().StringP("workdir", "C", "", "Change working directory before running command [default: crate root]")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"env":       env.ActionNameValues(false),
		"exec":      bridge.ActionCarapaceBin("cargo").Split(),
		"features":  cargo.ActionFeatures("").UniqueList(","),
		"shell":     bridge.ActionCarapaceBin().SplitP(),
		"use-shell": os.ActionShells(),
		"watch":     carapace.ActionFiles().List(","),
		"workdir":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			envs, err := cmd.Flags().GetStringArray("env")
			if err != nil {
				return carapace.ActionValues(err.Error())
			}
			for _, e := range envs {
				if k, v, ok := strings.Cut(e, "="); ok {
					c.Setenv(k, v)
				}
			}
			return action.Chdir(rootCmd.Flag("workdir").Value.String()).
				Invoke(c).ToA()
		})
	})
}
