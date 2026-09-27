package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a key from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_help_deleteCmd).Standalone()

	pkg_helpCmd.AddCommand(pkg_help_deleteCmd)
}
