package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{
	Use:   "lsp",
	Short: "Language server operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lspCmd).Standalone()

	lspCmd.Flags().String("cache-path", "", "Set a cache path")
	lspCmd.Flags().StringP("config", "c", "", "Path to the Taplo configuration file")
	lspCmd.Flags().Bool("no-auto-config", false, "Do not search for a configuration file")
	rootCmd.AddCommand(lspCmd)

	carapace.Gen(lspCmd).FlagCompletion(carapace.ActionMap{
		"cache-path": carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
	})
}
