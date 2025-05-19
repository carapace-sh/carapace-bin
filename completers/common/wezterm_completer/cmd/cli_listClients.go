package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cli_listClientsCmd = &cobra.Command{
	Use:   "list-clients [OPTIONS]",
	Short: "list clients",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_listClientsCmd).Standalone()

	cli_listClientsCmd.Flags().String("format", "", "Controls the output format")
	cli_listClientsCmd.Flags().BoolP("help", "h", false, "Print help")
	cliCmd.AddCommand(cli_listClientsCmd)

	carapace.Gen(cli_listClientsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})
}
