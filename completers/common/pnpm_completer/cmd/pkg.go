package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Manages your package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkgCmd).Standalone()

	pkgCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pkgCmd.PersistentFlags().Bool("json", false, "When setting, parse the value as JSON. When getting a single key, return its JSON-encoded form instead of the raw value")
	rootCmd.AddCommand(pkgCmd)
}
