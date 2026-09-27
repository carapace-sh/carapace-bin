package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_listRegistriesCmd = &cobra.Command{
	Use:   "list-registries",
	Short: "Lists all registries that have their metadata cache locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_listRegistriesCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_listRegistriesCmd)
}
