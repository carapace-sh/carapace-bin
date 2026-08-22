package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "returns the path to the specified app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prefixCmd).Standalone()
	rootCmd.AddCommand(prefixCmd)
}
