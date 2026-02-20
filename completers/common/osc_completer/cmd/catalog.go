package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Service catalog commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var catalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List service catalog entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalogCmd).Standalone()
	carapace.Gen(catalogListCmd).Standalone()

	rootCmd.AddCommand(catalogCmd)
	catalogCmd.AddCommand(catalogListCmd)
}
