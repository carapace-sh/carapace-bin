package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/fish"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fish",
	Short: "the friendly interactive shell",
	Long:  "https://fishshell.com/",
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

	rootCmd.Flags().StringP("command", "c", "", "Evaluate  the  specified  commands  instead  of  reading  from the commandline")
	rootCmd.Flags().StringP("debug", "d", "", "Enables debug output and specify a pattern for matching debug categories")
	rootCmd.Flags().StringP("debug-output", "o", "", "Specifies a file path to receive the debug output")
	rootCmd.Flags().StringP("features", "f", "", "Enables one or more comma-separated feature flags")
	rootCmd.Flags().StringP("init-command", "C", "", "Evaluate specified commands after reading the configuration")
	rootCmd.Flags().BoolP("interactive", "i", false, "The shell is interactive")
	rootCmd.Flags().BoolP("login", "l", false, "Act as if invoked as a login shell")
	rootCmd.Flags().BoolP("no-config", "N", false, "Do not read configuration files")
	rootCmd.Flags().BoolP("no-execute", "n", false, "Do not execute any commands, only perform syntax checking")
	rootCmd.Flags().Bool("print-debug-categories", false, "Print all debug categories, and then exit")
	rootCmd.Flags().Bool("print-rusage-self", false, "When fish exits, output stats from getrusage")
	rootCmd.Flags().BoolP("private", "P", false, "Enables private mode")
	rootCmd.Flags().StringP("profile", "p", "", "when fish exits, output timing information on all executed commands to the specified file")
	rootCmd.Flags().String("profile-startup", "", "Will write timing for fish startup to specified file")
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"command":      bridge.ActionCarapaceBin().SplitP(),
		"debug":        fish.ActionDebugCategories().UniqueList(","),
		"debug-output": carapace.ActionFiles(),
		"features": carapace.ActionValuesDescribed(
			"stderr-nocaret", "^ no longer redirects stderr",
			"qmark-noglob", "? no longer globs",
			"regex-easyesc", `string replace -r needs fewer \'s`,
		).UniqueList(","),
		"init-command":    bridge.ActionCarapaceBin().SplitP(),
		"profile":         carapace.ActionFiles(),
		"profile-startup": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("command").Changed {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)
}
