package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugwhyunstableCmd = &cobra.Command{
	Use:    "debugwhyunstable",
	Short:  "REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugwhyunstableCmd).Standalone()

	rootCmd.AddCommand(debugwhyunstableCmd)
}
