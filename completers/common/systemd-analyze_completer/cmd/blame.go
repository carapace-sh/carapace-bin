package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Print list of running units ordered by time to init",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	rootCmd.AddCommand(blameCmd)
}
