package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugdirstateignorepatternshashCmd = &cobra.Command{
	Use:    "debugdirstateignorepatternshash",
	Short:  "show the hash of ignore patterns stored in dirstate if v2,",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdirstateignorepatternshashCmd).Standalone()

	rootCmd.AddCommand(debugdirstateignorepatternshashCmd)
}
