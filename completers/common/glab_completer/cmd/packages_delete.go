package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packages_deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Short:   "Delete a package from a project's package registry.",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packages_deleteCmd).Standalone()

	packages_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	packagesCmd.AddCommand(packages_deleteCmd)
}
