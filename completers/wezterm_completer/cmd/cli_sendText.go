package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_sendTextCmd = &cobra.Command{
	Use:   "send-text [OPTIONS] [TEXT]",
	Short: "Send text to a pane as though it were pasted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_sendTextCmd).Standalone()

	cli_sendTextCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_sendTextCmd.Flags().Bool("no-paste", false, "Send the text directly, rather than as a bracketed paste")
	cli_sendTextCmd.Flags().String("pane-id", "", "Specify the target pane")
	cliCmd.AddCommand(cli_sendTextCmd)

	carapace.Gen(cli_sendTextCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})
}
