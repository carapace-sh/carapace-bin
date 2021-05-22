package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Update configuration with a value for the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	config_setCmd.Flags().StringP("host", "h", "", "Set per-host setting")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})

	carapace.Gen(config_setCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"git_protocol", "What protocol to use when performing git operations.",
			"editor", "What editor gh should run when creating issues, pull requests, etc.",
			"prompt", "toggle interactive prompting in the terminal",
			"pager", "the terminal pager program to send standard output to",
		),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "git_protocol":
				return carapace.ActionValues("ssh", "https")
			case "editor":
				return carapace.ActionValues("emacs", "micro", "nano", "nvim", "vi", "vim")
			case "prompt":
				return carapace.ActionValues("enabled", "disabled")
			case "pager":
				return carapace.ActionValues("bat --style grid", "more", "most", "less")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
