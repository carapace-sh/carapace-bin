package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Manages your package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pkgCmd).Standalone()

	helpCmd.AddCommand(help_pkgCmd)
}
