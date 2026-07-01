package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plistCmd = &cobra.Command{
	Use:   "plist",
	Short: "Prints a property list embedded in a binary",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(plistCmd).Standalone()
	rootCmd.AddCommand(plistCmd)
	carapace.Gen(plistCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
