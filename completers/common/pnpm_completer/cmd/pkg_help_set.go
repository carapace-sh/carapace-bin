package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a value in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_help_setCmd).Standalone()

	pkg_helpCmd.AddCommand(pkg_help_setCmd)
}
