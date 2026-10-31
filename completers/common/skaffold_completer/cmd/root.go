package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/skaffold_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skaffold",
	Short: "A tool that facilitates continuous development for Kubernetes applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("color", "34", "Specify the default output color in ANSI escape codes")
	rootCmd.PersistentFlags().Bool("force-colors", false, "Always print color codes (hidden)")
	rootCmd.PersistentFlags().Bool("interactive", true, "Allow user prompts for more information")
	rootCmd.PersistentFlags().Bool("timestamps", false, "Print timestamps in logs")
	rootCmd.PersistentFlags().Bool("update-check", true, "Check for a more recent version of Skaffold")
	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Log level: one of [panic fatal error warning info debug trace]")
	rootCmd.Flag("force-colors").Hidden = true
}

// applyCommonFlagCompletion registers completion for flags shared by most skaffold commands
func applyCommonFlagCompletion(cmd *cobra.Command) {
	actions := carapace.ActionMap{
		"cache-file":     carapace.ActionFiles(),
		"config":         carapace.ActionFiles(),
		"event-log-file": carapace.ActionFiles(),
		"file-output":    carapace.ActionFiles(),
		"filename":       carapace.ActionFiles(".yaml", ".yml"),
		"kube-context":   kubectl.ActionContexts(),
		"kubeconfig":     carapace.ActionFiles(),
		"last-log-file":  carapace.ActionFiles(),
		"mute-logs":      carapace.ActionValues("all", "build", "deploy", "none", "status-check"),
		"profile":        action.ActionProfiles(cmd),
		"sync-remote-cache": carapace.ActionValues(
			"always",
			"missing",
			"never",
		),
		"verbosity": carapace.ActionValues("debug", "error", "fatal", "info", "panic", "trace", "warning"),
	}

	available := make(carapace.ActionMap)
	for name, a := range actions {
		if cmd.LocalFlags().Lookup(name) != nil {
			available[name] = a
		}
	}
	carapace.Gen(cmd).FlagCompletion(available)
}
