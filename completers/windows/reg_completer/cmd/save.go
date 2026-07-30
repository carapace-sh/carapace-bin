package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save a subkey to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(saveCmd).Standalone()
	rootCmd.AddCommand(saveCmd)
}
