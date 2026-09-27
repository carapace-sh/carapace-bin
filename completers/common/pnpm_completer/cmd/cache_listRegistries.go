package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_listRegistriesCmd = &cobra.Command{
	Use:   "list-registries",
	Short: "Lists all registries that have their metadata cache locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listRegistriesCmd).Standalone()

	cache_listRegistriesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cacheCmd.AddCommand(cache_listRegistriesCmd)
}
