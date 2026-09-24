package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Safely remove `node_modules` directories from the current project (or every workspace project) without following NTFS junctions into their targets. A `clean` script in `package.json` overrides the built-in command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cleanCmd.Flags().BoolP("lockfile", "l", false, "Also remove `pnpm-lock.yaml` files")
	rootCmd.AddCommand(cleanCmd)
}
