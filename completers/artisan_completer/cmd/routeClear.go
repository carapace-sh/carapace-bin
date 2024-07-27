package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var routeClearCmd = &cobra.Command{
	Use:   "route:clear",
	Short: "Remove the route cache file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routeClearCmd).Standalone()

	rootCmd.AddCommand(routeClearCmd)
}
