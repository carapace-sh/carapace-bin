package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(InfoCmd).Standalone()
	rootCmd.AddCommand(InfoCmd)
}
