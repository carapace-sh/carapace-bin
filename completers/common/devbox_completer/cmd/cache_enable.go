package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable the Devbox Nix cache for your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_enableCmd).Standalone()

	cacheCmd.AddCommand(cache_enableCmd)
}
