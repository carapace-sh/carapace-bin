package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var show_help_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Show the default profile used for the `rustup install` command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_help_profileCmd).Standalone()

	show_helpCmd.AddCommand(show_help_profileCmd)
}
