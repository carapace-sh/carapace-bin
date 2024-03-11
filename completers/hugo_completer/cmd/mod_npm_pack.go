package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_npm_packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Experimental: Prepares and writes a composite package.json file for your project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_npm_packCmd).Standalone()
	mod_npmCmd.AddCommand(mod_npm_packCmd)
}
