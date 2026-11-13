package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ruff",
	Short: "Ruff: An extremely fast Python linter and code formatter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("color", "", "Control when colored output is used")
	rootCmd.PersistentFlags().StringSlice("config", nil, "Either a path to a TOML configuration file (`pyproject.toml` or `ruff.toml`), or a TOML `<KEY> = <VALUE>` pair (such as you might find in a `ruff.toml` configuration file) overriding a specific configuration option (e.g., `--config \"line-length = 100\"` or `--config \"format.quote-style = 'single'\"`). Overrides of individual settings using this option always take precedence over all configuration files, including configuration files that were also specified using `--config`")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("isolated", false, "Ignore all configuration files")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Print diagnostics, but nothing else")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "Disable all logging (but still exit with status code \"1\" upon detecting diagnostics)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
