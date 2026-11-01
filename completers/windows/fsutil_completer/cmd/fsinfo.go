package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsinfoCmd = &cobra.Command{
	Use:   "fsinfo",
	Short: "file system information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsinfoCmd).Standalone()
	rootCmd.AddCommand(fsinfoCmd)
}
