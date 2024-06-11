package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tag_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_helpCmd).Standalone()

	tagCmd.AddCommand(tag_helpCmd)
}
