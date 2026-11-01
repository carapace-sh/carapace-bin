package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Update configuration with a value for the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().StringP("host", "h", "", "Set per-host setting")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})

	carapace.Gen(config_setCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"browser", "the web browser to use for opening URLs",
			"editor", "What editor gh should run when creating issues, pull requests, etc.",
			"git_protocol", "What protocol to use when performing git operations.",
			"http_unix_socket", "the path to a unix socket through which to make HTTP connection",
			"pager", "the terminal pager program to send standard output to",
			"prompt", "toggle interactive prompting in the terminal",
		),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "browser":
				return carapace.ActionFiles()
			case "git_protocol":
				return carapace.ActionValues("https", "ssh")
			case "editor":
				return carapace.ActionValues("emacs", "micro", "nano", "nvim", "vi", "vim")
			case "prompt":
				return carapace.ActionValues("disabled", "enabled").StyleF(style.ForKeyword)
			case "pager":
				return carapace.ActionValues("bat --style grid", "less", "more", "most")
			case "http_unix_socket":
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
