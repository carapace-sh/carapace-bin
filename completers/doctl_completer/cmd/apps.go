package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Display commands for working with apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsCmd).Standalone()
	rootCmd.AddCommand(appsCmd)
}
