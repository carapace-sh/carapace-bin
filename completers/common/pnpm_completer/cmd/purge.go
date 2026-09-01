package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Alias of `clean`: same behavior, except a `purge` script (not a `clean` script) overrides it when present",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	purgeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	purgeCmd.Flags().BoolP("lockfile", "l", false, "Also remove `pnpm-lock.yaml` files")
	rootCmd.AddCommand(purgeCmd)
}
