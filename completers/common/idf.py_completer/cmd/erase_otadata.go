package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eraseOtadataCmd = &cobra.Command{
	Use:   "erase-otadata",
	Short: "Erase otadata partition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eraseOtadataCmd).Standalone()
	rootCmd.AddCommand(eraseOtadataCmd)
}
