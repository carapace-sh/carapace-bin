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
	Long:  "https://fishshell.com/docs/current/cmds/fish.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("command", "c", "", "evaluate specified commands")
	rootCmd.Flags().StringP("init-command", "C", "", "evaluate commands after config")
	rootCmd.Flags().StringP("debug", "d", "", "enable debug output")
	rootCmd.Flags().StringP("debug-output", "o", "", "debug output file")
	rootCmd.Flags().StringP("features", "f", "", "enable feature flags")
	rootCmd.Flags().BoolP("help", "h", false, "display help")
	rootCmd.Flags().BoolP("interactive", "i", false, "interactive shell")
	rootCmd.Flags().BoolP("login", "l", false, "login shell")
	rootCmd.Flags().BoolP("no-config", "N", false, "do not read config files")
	rootCmd.Flags().BoolP("no-execute", "n", false, "only syntax checking")
	rootCmd.Flags().BoolP("private", "P", false, "private mode")
	rootCmd.Flags().StringP("profile", "p", "", "output timing info")
	rootCmd.Flags().String("profile-startup", "", "write startup timing")
	rootCmd.Flags().Bool("print-rusage-self", false, "output getrusage stats")
	rootCmd.Flags().Bool("print-debug-categories", false, "print debug categories")
	rootCmd.Flags().BoolP("version", "v", false, "print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"command":         bridge.ActionCarapaceBin().SplitP(),
		"debug":           fish.ActionDebugCategories().UniqueList(","),
		"debug-output":    carapace.ActionFiles(),
		"features":        carapace.ActionValues("stderr-nocaret", "qmark-noglob", "regex-easyesc").UniqueList(","),
		"init-command":    bridge.ActionCarapaceBin().SplitP(),
		"profile":         carapace.ActionFiles(),
		"profile-startup": carapace.ActionFiles(),
	})
}
