package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_set_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "The default components installed with a toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_set_profileCmd).Standalone()

	help_setCmd.AddCommand(help_set_profileCmd)
}
