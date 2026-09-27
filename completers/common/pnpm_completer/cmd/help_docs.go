package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Opens the documentation of a package in the browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_docsCmd).Standalone()

	helpCmd.AddCommand(help_docsCmd)
}
