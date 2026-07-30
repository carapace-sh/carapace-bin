package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glCmd = &cobra.Command{
	Use:   "gl",
	Short: "get event log configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glCmd).Standalone()
	rootCmd.AddCommand(glCmd)
}
