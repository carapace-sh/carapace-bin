package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scripts_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_help_getCmd).Standalone()

	scripts_helpCmd.AddCommand(scripts_help_getCmd)
}
