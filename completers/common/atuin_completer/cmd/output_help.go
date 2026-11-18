package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var output_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(output_helpCmd).Standalone()

	outputCmd.AddCommand(output_helpCmd)
}
