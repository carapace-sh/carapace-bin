package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_scripts_getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_scripts_getCmd).Standalone()

	help_scriptsCmd.AddCommand(help_scripts_getCmd)
}
