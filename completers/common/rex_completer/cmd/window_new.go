package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var windowNewCmd = &cobra.Command{
	Use:   "new [LABEL] [-- COMMAND...]",
	Short: "Create a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowNewCmd).Standalone()

	windowNewCmd.Flags().SetInterspersed(false)

	windowNewCmd.Flags().StringP("cwd", "c", "", "working directory for the terminal")
	windowNewCmd.Flags().Bool("focus", true, "make the new window active")
	windowNewCmd.Flags().Bool("json", false, "output the created window as JSON")
	windowNewCmd.Flags().Bool("keep-open", false, "keep the terminal open after its command exits")
	windowNewCmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	windowNewCmd.Flags().String("shell", "", "run COMMAND through SHELL's login environment")

	windowCmd.AddCommand(windowNewCmd)

	carapace.Gen(windowNewCmd).FlagCompletion(carapace.ActionMap{
		"cwd":     carapace.ActionDirectories(),
		"session": action.ActionSessions(windowNewCmd),
		"shell":   carapace.ActionValues("none"),
	})

	carapace.Gen(windowNewCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
