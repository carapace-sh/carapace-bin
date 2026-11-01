package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print cmd/tailscale environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_envCmd).Standalone()

	debugCmd.AddCommand(debug_envCmd)
}
