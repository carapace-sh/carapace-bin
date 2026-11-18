package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help___internal_prepareSearchIndexCmd = &cobra.Command{
	Use:   "prepare-search-index",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help___internal_prepareSearchIndexCmd).Standalone()

	help___internalCmd.AddCommand(help___internal_prepareSearchIndexCmd)
}
