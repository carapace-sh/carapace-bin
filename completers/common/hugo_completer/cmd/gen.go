package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A collection of several useful generators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genCmd).Standalone()
	rootCmd.AddCommand(genCmd)
}
