package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"view"},
	Short:   "Shows details of a configuration",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_showCmd).Standalone()

	configureCmd.AddCommand(configure_showCmd)
}
