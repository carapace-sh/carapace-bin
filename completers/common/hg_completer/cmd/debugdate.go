package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugdateCmd = &cobra.Command{
	Use:    "debugdate",
	Short:  "parse and display a date",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdateCmd).Standalone()

	debugdateCmd.Flags().BoolP("extended", "e", false, "try extended date formats")
	rootCmd.AddCommand(debugdateCmd)
}
