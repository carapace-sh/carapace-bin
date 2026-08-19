package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listPluginsCmd = &cobra.Command{
	Use:   "list-plugins",
	Short: "list available nixops plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPluginsCmd).Standalone()
	rootCmd.AddCommand(listPluginsCmd)
}
