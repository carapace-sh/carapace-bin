package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Generates a pnpm-lock.yaml from an external lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	importCmd.Flags().String("pnpr-server", "", "URL of a pnpr server. Accepted for symmetry with the other installing commands; `pnpm import` always resolves locally")
	rootCmd.AddCommand(importCmd)
}
