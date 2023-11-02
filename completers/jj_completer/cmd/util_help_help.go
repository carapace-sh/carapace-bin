package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var util_help_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_help_helpCmd).Standalone()

	util_helpCmd.AddCommand(util_help_helpCmd)
}
