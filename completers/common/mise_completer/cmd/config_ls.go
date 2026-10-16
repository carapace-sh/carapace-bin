package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List config files currently in use",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_lsCmd).Standalone()

	configCmd.AddCommand(config_lsCmd)
}
