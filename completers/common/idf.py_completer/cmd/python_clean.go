package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pythonCleanCmd = &cobra.Command{
	Use:   "python-clean",
	Short: "Delete generated Python byte code from the IDF directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pythonCleanCmd).Standalone()

	rootCmd.AddCommand(pythonCleanCmd)
}
