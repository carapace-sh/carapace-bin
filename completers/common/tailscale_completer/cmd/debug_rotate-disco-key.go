package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_rotate_disco_keyCmd = &cobra.Command{
	Use:   "rotate-disco-key",
	Short: "Rotate the discovery key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_rotate_disco_keyCmd).Standalone()

	debugCmd.AddCommand(debug_rotate_disco_keyCmd)
}
