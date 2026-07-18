package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Manage package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkgCmd).Standalone()
	pkgCmd.PersistentFlags().BoolP("force", "f", false, "remove various protections against unfortunate side effects")
	pkgCmd.PersistentFlags().Bool("json", false, "output as json")
	addWorkspaceFlags(pkgCmd)

	rootCmd.AddCommand(pkgCmd)
}
