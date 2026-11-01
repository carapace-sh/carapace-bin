package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frontCmd = &cobra.Command{
	Use:   "front",
	Short: "Show the front application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frontCmd).Standalone()
	rootCmd.AddCommand(frontCmd)
}
