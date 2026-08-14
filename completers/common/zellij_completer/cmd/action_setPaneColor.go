package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_setPaneColorCmd = &cobra.Command{
	Use:   "set-pane-color",
	Short: "Set the default foreground/background color of a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_setPaneColorCmd).Standalone()

	action_setPaneColorCmd.Flags().String("bg", "", "Background color (e.g. \"#001a3a\", \"rgb:00/1a/3a\")")
	action_setPaneColorCmd.Flags().String("fg", "", "Foreground color (e.g. \"#00e000\", \"rgb:00/e0/00\")")
	action_setPaneColorCmd.Flags().BoolP("help", "h", false, "Print help")
	action_setPaneColorCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3). Defaults to $ZELLIJ_PANE_ID if not provided")
	action_setPaneColorCmd.Flags().Bool("reset", false, "Reset pane colors to terminal defaults")
	actionCmd.AddCommand(action_setPaneColorCmd)
}
