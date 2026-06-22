package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var assetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Asset pipeline tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(assetsCmd).Standalone()
	assetsCmd.Flags().BoolP("help", "h", false, "Show help")
}
