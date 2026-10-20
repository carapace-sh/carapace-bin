// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Install a tool and add it to configuration",
	Aliases: []string{"u"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()

	useCmd.Flags().String("before", "", "Only install versions released before this date or older than this duration")
	useCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	useCmd.Flags().Bool("dry-run-code", false, "Like --dry-run, but exit with code 1 if there are changes to make")
	useCmd.Flags().StringP("env", "e", "", "Write mise.<ENV>.toml in the current directory, or .mise.<ENV>.toml if that exists")
	useCmd.Flags().BoolP("force", "f", false, "Reinstall even if already installed")
	useCmd.Flags().Bool("fuzzy", false, "Write the version as requested, such as `20` for node@20")
	useCmd.Flags().BoolP("global", "g", false, "Write to the global config instead of the project config")
	useCmd.Flags().BoolP("help", "h", false, "Print help")
	useCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel (default: the `jobs` setting)")
	useCmd.Flags().String("minimum-release-age", "", "Only install versions released before this date or older than this duration")
	useCmd.Flags().Bool("override", false, "Write to the global override file, `config.override.toml`")
	useCmd.Flags().StringP("path", "p", "", "Write to this config file, or to the config file in this directory")
	useCmd.Flags().Bool("pin", false, "Write the resolved version, such as `20.19.5` for node@20")
	useCmd.Flags().Bool("raw", false, "Connect the install commands' stdin, stdout, and stderr to the terminal")
	useCmd.Flags().StringSlice("remove", nil, "Remove these tools from the config file")
	useCmd.Flags().StringSlice("rm", nil, "Remove these tools from the config file")
	useCmd.Flags().StringSlice("unset", nil, "Remove these tools from the config file")
	useCmd.Flag("before").Hidden = true
	useCmd.Flag("rm").Hidden = true
	useCmd.Flag("unset").Hidden = true
	rootCmd.AddCommand(useCmd)

	carapace.Gen(useCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})

	carapace.Gen(useCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
