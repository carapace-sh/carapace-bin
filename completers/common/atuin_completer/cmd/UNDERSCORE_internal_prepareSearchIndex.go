package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var __internal_prepareSearchIndexCmd = &cobra.Command{
	Use:   "prepare-search-index",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(__internal_prepareSearchIndexCmd).Standalone()

	__internalCmd.AddCommand(__internal_prepareSearchIndexCmd)
}
