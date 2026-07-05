package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tieringCmd = &cobra.Command{
	Use:   "tiering",
	Short: "storage tier management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tieringCmd).Standalone()
	rootCmd.AddCommand(tieringCmd)
}
