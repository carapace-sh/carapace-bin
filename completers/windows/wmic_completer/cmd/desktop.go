package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopCmd = &cobra.Command{
	Use:   "desktop",
	Short: "desktop management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopCmd).Standalone()
	rootCmd.AddCommand(desktopCmd)
}
