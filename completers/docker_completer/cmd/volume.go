package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:     "volume COMMAND",
	Short:   "Manage volumes",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volumeCmd).Standalone()

	rootCmd.AddCommand(volumeCmd)
}
