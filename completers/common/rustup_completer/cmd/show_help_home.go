package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var show_help_homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Display the computed value of RUSTUP_HOME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_help_homeCmd).Standalone()

	show_helpCmd.AddCommand(show_help_homeCmd)
}
