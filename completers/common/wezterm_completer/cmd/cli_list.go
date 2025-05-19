package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cli_listCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "list windows, tabs and panes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_listCmd).Standalone()

	cli_listCmd.Flags().String("format", "", "Controls the output format")
	cli_listCmd.Flags().BoolP("help", "h", false, "Print help")
	cliCmd.AddCommand(cli_listCmd)

	carapace.Gen(cli_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})
}
