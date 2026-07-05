package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var distfilesCmd = &cobra.Command{
	Use:   "distfiles",
	Short: "Display each distfile, its checksums, and fetch URLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distfilesCmd).Standalone()
	rootCmd.AddCommand(distfilesCmd)
}
