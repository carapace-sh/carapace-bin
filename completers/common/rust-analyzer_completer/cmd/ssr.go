package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssrCmd = &cobra.Command{
	Use:   "ssr",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssrCmd).Standalone()

	rootCmd.AddCommand(ssrCmd)
}
