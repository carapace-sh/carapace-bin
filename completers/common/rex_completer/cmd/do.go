package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do (ACTION | SCRIPT | -e CODE) [KEY=VALUE...]",
	Short: "Perform an action by name, or run a Lua script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doCmd).Standalone()

	doCmd.Flags().StringP("action", "a", "", "run the action NAME that the script defines")
	doCmd.Flags().Bool("all", false, "listen for events from every session")
	doCmd.Flags().String("args", "", "pass the arguments as a JSON object (@FILE, or - for stdin)")
	doCmd.Flags().StringP("block", "b", "", "apply the action to BLOCK")
	doCmd.Flags().StringP("eval", "e", "", "run CODE as a Lua script")
	doCmd.Flags().Bool("fail-fast", false, "exit at the first rex.on handler that fails")
	doCmd.Flags().String("for", "", "stop listening for events after DURATION")
	doCmd.Flags().Bool("json", false, "print the result as JSON")
	doCmd.Flags().String("origin", "cli", "report the origin as ORIGIN: key, palette, api, or cli")
	doCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	rootCmd.AddCommand(doCmd)

	carapace.Gen(doCmd).FlagCompletion(carapace.ActionMap{
		"action":  action.ActionClientActions(doCmd),
		"args":    carapace.ActionFiles().Prefix("@"),
		"block":   action.ActionBlocks(doCmd),
		"for":     carapace.ActionValues("1s", "10s", "1m", "10m"),
		"origin":  carapace.ActionValues("api", "cli", "key", "palette"),
		"session": action.ActionSessions(doCmd),
	})

	carapace.Gen(doCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionClientActions(doCmd),
			carapace.ActionValues("-"),
		).ToA(),
	)

	carapace.Gen(doCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionClientActionArgs(doCmd).Suffix("=")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
