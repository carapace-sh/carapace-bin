package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_scripts_newCmd = &cobra.Command{
	Use:   "new",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_scripts_newCmd).Standalone()

	help_scriptsCmd.AddCommand(help_scripts_newCmd)
}
