package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entity_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for New Relic entities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entity_searchCmd).Standalone()
	entity_searchCmd.Flags().StringP("alert-severity", "s", "", "search for entities matching the given alert severity type")
	entity_searchCmd.Flags().StringP("domain", "d", "", "search for entities matching the given entity domain")
	entity_searchCmd.Flags().StringSliceP("fields-filter", "f", nil, "filter search results to only return certain fields for each search result")
	entity_searchCmd.Flags().StringP("name", "n", "", "search for entities matching the given name")
	entity_searchCmd.Flags().StringP("reporting", "r", "", "search for entities based on whether or not an entity is reporting (true or false)")
	entity_searchCmd.Flags().String("tag", "", "search for entities matching the given entity tag")
	entity_searchCmd.Flags().StringP("type", "t", "", "search for entities matching the given type")
	entityCmd.AddCommand(entity_searchCmd)

	// TODO flag completion
	carapace.Gen(entity_searchCmd).FlagCompletion(carapace.ActionMap{
		"reporting": carapace.ActionValues("true", "false"),
	})
}
