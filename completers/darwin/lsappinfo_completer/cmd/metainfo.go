package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metainfoCmd = &cobra.Command{
	Use:   "metainfo",
	Short: "Show meta information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metainfoCmd).Standalone()
	rootCmd.AddCommand(metainfoCmd)
}
