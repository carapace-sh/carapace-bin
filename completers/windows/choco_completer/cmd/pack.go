package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "pack a nuspec into a nupkg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()
	rootCmd.AddCommand(packCmd)
}
