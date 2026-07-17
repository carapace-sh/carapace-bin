package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpjpcategoryCmd = &cobra.Command{
	Use:   "dumpjpcategory",
	Short: "Dump the jetsam properties category for all services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpjpcategoryCmd).Standalone()
	rootCmd.AddCommand(dumpjpcategoryCmd)
}
