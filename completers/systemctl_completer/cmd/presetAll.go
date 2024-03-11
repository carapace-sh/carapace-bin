package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var presetAllCmd = &cobra.Command{
	Use:     "preset-all",
	Short:   "Enable/disable all unit files based on preset configuration",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(presetAllCmd).Standalone()

	rootCmd.AddCommand(presetAllCmd)
}
