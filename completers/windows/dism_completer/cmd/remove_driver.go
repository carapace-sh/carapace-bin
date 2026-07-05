package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RemoveDriverCmd = &cobra.Command{
	Use:   "Remove-Driver",
	Short: "remove third-party drivers from an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RemoveDriverCmd).Standalone()
	rootCmd.AddCommand(RemoveDriverCmd)
}
