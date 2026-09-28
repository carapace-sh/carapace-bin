package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugformatCmd = &cobra.Command{
	Use:    "debugformat",
	Short:  "display format information about the current repository",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugformatCmd).Standalone()

	debugformatCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugformatCmd)
}
