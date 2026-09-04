package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efCmd = &cobra.Command{
	Use:   "ef",
	Short: "Entity Framework Core tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efCmd).Standalone()
	rootCmd.AddCommand(efCmd)
}
