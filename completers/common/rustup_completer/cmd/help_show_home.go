package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_show_homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Display the computed value of RUSTUP_HOME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_show_homeCmd).Standalone()

	help_showCmd.AddCommand(help_show_homeCmd)
}
