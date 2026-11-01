package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Package a kcl package into tar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_pkgCmd).Standalone()

	mod_pkgCmd.Flags().String("target", "", "packaged target path")
	mod_pkgCmd.Flags().Bool("vendor", false, "package in vendor mode (default: false)")
	modCmd.AddCommand(mod_pkgCmd)

	carapace.Gen(mod_pkgCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionDirectories(),
	})
}
