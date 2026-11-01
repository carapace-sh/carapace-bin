package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

func AddCreateFlags(cmd *cobra.Command, splitDefault string) {
	cmd.Flags().StringP("block", "b", "", "use BLOCK as the split anchor")
	cmd.Flags().StringP("cwd", "c", "", "working directory for the terminal")
	cmd.Flags().Bool("focus", true, "focus the new block")
	cmd.Flags().Bool("json", false, "output the creation result as JSON")
	cmd.Flags().Bool("keep-open", false, "keep the terminal open after its command exits")
	cmd.Flags().String("label", "", "set the new block label to TEXT")
	cmd.Flags().String("ratio", "", "give the new block PCT percent of the split")
	cmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	cmd.Flags().String("shell", "", "run COMMAND through SHELL's login environment")
	cmd.Flags().String("split", splitDefault, "place the new block in DIRECTION relative to the anchor")
	cmd.Flags().Bool("wait", false, "wait and return the terminal child's exit status")
	cmd.Flags().StringP("window", "w", "", "scope to window WINDOW")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"block":   action.ActionBlocks(cmd),
		"cwd":     carapace.ActionDirectories(),
		"session": action.ActionSessions(cmd),
		"shell":   carapace.ActionValues("none"),
		"split":   rex.ActionDirections(),
		"window":  action.ActionWindows(cmd),
	})
}

func AddSessionCreateFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cwd", "c", "", "working directory for the terminal")
	cmd.Flags().Bool("json", false, "output the created session as JSON")
	cmd.Flags().Bool("keep-open", false, "keep the terminal open after its command exits")
	cmd.Flags().String("shell", "", "run COMMAND through SHELL's login environment")
	cmd.Flags().String("window", "", "label for the first window")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"cwd":   carapace.ActionDirectories(),
		"shell": carapace.ActionValues("none"),
	})
}
