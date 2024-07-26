package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaCardCmd = &cobra.Command{
	Use:   "nova:card <name>",
	Short: "Create a new card",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaCardCmd).Standalone()

	rootCmd.AddCommand(novaCardCmd)
}
