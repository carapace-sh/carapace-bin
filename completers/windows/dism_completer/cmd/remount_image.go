package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RemountImageCmd = &cobra.Command{
	Use:   "Remount-Image",
	Short: "remount an inaccessible mounted image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RemountImageCmd).Standalone()
	rootCmd.AddCommand(RemountImageCmd)
}
