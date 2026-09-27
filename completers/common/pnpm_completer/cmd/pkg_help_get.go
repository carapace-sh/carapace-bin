package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a value from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_help_getCmd).Standalone()

	pkg_helpCmd.AddCommand(pkg_help_getCmd)
}
