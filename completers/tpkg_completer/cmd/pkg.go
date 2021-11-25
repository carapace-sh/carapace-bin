package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Manage packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkgCmd).Standalone()
	pkgCmd.PersistentFlags().String("project-root", "", "Specify the project root")
	rootCmd.AddCommand(pkgCmd)

	carapace.Gen(pkgCmd).FlagCompletion(carapace.ActionMap{
		"project-root": carapace.ActionDirectories(),
	})
}
