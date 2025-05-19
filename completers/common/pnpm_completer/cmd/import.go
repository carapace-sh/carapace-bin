package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Short:   "Generates pnpm-lock.yaml from an npm package-lock.json",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	rootCmd.AddCommand(importCmd)
}
