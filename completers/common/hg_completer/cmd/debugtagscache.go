package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugtagscacheCmd = &cobra.Command{
	Use:    "debugtagscache",
	Short:  "display the contents of .hg/cache/hgtagsfnodes1",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugtagscacheCmd).Standalone()

	rootCmd.AddCommand(debugtagscacheCmd)
}
