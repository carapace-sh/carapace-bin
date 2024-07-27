package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskPurgeCmd = &cobra.Command{
	Use:   "dusk:purge",
	Short: "Purge dusk test debugging files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskPurgeCmd).Standalone()

	rootCmd.AddCommand(duskPurgeCmd)
}
