package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavorCmd).Standalone()

	rootCmd.AddCommand(flavorCmd)
}
