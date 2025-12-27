package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chidCmd = &cobra.Command{
	Use:   "chid",
	Short: "List local CHIDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chidCmd).Standalone()

	rootCmd.AddCommand(chidCmd)
}
