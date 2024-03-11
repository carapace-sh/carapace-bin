package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "benthos",
	Short: "A stream processor for mundane tasks",
	Long:  "https://www.benthos.dev",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("chilled", false, "continue to execute a config containing linter errors")
	rootCmd.Flags().StringP("config", "c", "", "a path to a configuration file")
	rootCmd.Flags().StringP("env-file", "e", "", "import environment variables from a dotenv file")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().String("log.level", "", "override the configured log level, options are: off, error, warn, info, debug, trace")
	rootCmd.Flags().StringP("resources", "r", "", "pull in extra resources from a file, which can be referenced the same as resources defined in the main config, supports glob patterns (requires quotes)")
	rootCmd.Flags().StringP("set", "s", "", "set a field (identified by a dot path) in the main configuration file, e.g. `\"metrics.type=prometheus\"`")
	rootCmd.Flags().StringP("templates", "t", "", "EXPERIMENTAL: import Benthos templates, supports glob patterns (requires quotes)")
	rootCmd.Flags().BoolP("version", "v", false, "display version info, then exit")
	rootCmd.Flags().BoolP("watcher", "w", false, "EXPERIMENTAL: watch config files for changes and automatically apply them")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":    carapace.ActionFiles(),
		"env-file":  carapace.ActionFiles(),
		"log.level": carapace.ActionValues("off", "error", "warn", "info", "debug", "trace").StyleF(style.ForLogLevel),
		"resources": carapace.ActionValues(), // TODO
		"set":       carapace.ActionValues(), // TODO
		"templates": carapace.ActionValues(), // TODO
	})
}
