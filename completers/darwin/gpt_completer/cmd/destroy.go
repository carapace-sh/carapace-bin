package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy the GPT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()
	rootCmd.AddCommand(destroyCmd)

	destroyCmd.Flags().BoolP("recursive", "r", false, "Destroy recursively")
}