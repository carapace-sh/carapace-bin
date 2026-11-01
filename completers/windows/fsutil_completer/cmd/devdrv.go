package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devdrvCmd = &cobra.Command{
	Use:   "devdrv",
	Short: "manage developer drives",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devdrvCmd).Standalone()
	rootCmd.AddCommand(devdrvCmd)
}
