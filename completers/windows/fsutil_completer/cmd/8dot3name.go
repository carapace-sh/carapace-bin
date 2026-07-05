package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eightdot3nameCmd = &cobra.Command{
	Use:   "8dot3name",
	Short: "manage 8.3 short names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eightdot3nameCmd).Standalone()
	rootCmd.AddCommand(eightdot3nameCmd)
}
