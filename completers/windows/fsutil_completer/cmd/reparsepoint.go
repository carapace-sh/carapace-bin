package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reparsepointCmd = &cobra.Command{
	Use:   "reparsepoint",
	Short: "reparse point management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reparsepointCmd).Standalone()
	rootCmd.AddCommand(reparsepointCmd)
}
