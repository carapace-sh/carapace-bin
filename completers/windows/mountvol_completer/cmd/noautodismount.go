package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var noautodismountCmd = &cobra.Command{
	Use:   "noautodismount",
	Short: "disable automatic mounting of new volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(noautodismountCmd).Standalone()
	rootCmd.AddCommand(noautodismountCmd)
}
