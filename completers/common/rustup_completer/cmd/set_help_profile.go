package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_help_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "The default components installed with a toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_help_profileCmd).Standalone()

	set_helpCmd.AddCommand(set_help_profileCmd)
}
