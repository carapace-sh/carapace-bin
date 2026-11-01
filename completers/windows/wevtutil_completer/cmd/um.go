package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var umCmd = &cobra.Command{
	Use:   "um",
	Short: "uninstall event publisher manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(umCmd).Standalone()
	rootCmd.AddCommand(umCmd)
}
