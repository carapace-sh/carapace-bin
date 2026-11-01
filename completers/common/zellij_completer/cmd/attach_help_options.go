package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attach_help_optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Change the behaviour of zellij",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attach_help_optionsCmd).Standalone()

	attach_helpCmd.AddCommand(attach_help_optionsCmd)
}
