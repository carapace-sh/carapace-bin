package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugextensionsCmd = &cobra.Command{
	Use:    "debugextensions",
	Short:  "show information about active extensions",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugextensionsCmd).Standalone()

	debugextensionsCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugextensionsCmd)
}
