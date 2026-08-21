package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cachedCmd = &cobra.Command{
	Use:   "cached",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cachedCmd).Standalone()

	rootCmd.AddCommand(cachedCmd)
}
