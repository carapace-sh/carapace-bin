package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help for a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().String("format", "", "The output format (txt, xml, json, or md)")
	helpCmd.Flags().Bool("raw", false, "To output raw command help")

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "md", "txt", "xml"),
	})

	carapace.Gen(helpCmd).PositionalCompletion(
		action.ActionCommands(),
	)
}
