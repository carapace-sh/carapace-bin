package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "find the path to an installed executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()
	rootCmd.AddCommand(whichCmd)
}
