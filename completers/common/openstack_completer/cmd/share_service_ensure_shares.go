package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_service_ensure_sharesCmd = &cobra.Command{
	Use:   "shares",
	Short: "Run ensure shares in a back end (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_service_ensure_sharesCmd).Standalone()

	share_service_ensureCmd.AddCommand(share_service_ensure_sharesCmd)
}
