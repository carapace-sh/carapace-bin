package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "module maintenance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modCmd).Standalone()
	modCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(modCmd)
}
