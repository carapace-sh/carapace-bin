package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_restunCmd = &cobra.Command{
	Use:   "restun",
	Short: "Force a magicsock restun",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_restunCmd).Standalone()

	debugCmd.AddCommand(debug_restunCmd)
}
