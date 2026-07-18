package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installScripts_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List install-script approvals",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installScripts_lsCmd).Standalone()

	installScriptsCmd.AddCommand(installScripts_lsCmd)
}
