package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_show_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Show the default profile used for the `rustup install` command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_show_profileCmd).Standalone()

	help_showCmd.AddCommand(help_show_profileCmd)
}
