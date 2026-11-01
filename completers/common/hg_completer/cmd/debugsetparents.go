package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugsetparentsCmd = &cobra.Command{
	Use:    "debugsetparents",
	Short:  "REV1 [REV2]",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugsetparentsCmd).Standalone()

	rootCmd.AddCommand(debugsetparentsCmd)
}
