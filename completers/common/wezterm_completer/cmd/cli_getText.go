package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_getTextCmd = &cobra.Command{
	Use:   "get-text [OPTIONS]",
	Short: "Retrieves the textual content of a pane and output it to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_getTextCmd).Standalone()

	cli_getTextCmd.Flags().String("end-line", "", "The ending line number")
	cli_getTextCmd.Flags().Bool("escapes", false, "Include escape sequences that color and style the text")
	cli_getTextCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_getTextCmd.Flags().String("pane-id", "", "Specify the target pane")
	cli_getTextCmd.Flags().String("start-line", "", "The starting line number")
	cliCmd.AddCommand(cli_getTextCmd)

	carapace.Gen(cli_getTextCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})
}
