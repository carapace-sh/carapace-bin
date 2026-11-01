package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_attach_optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Change the behaviour of zellij",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_attach_optionsCmd).Standalone()

	help_attachCmd.AddCommand(help_attach_optionsCmd)
}
