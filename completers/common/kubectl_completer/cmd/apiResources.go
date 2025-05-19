package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var apiResourcesCmd = &cobra.Command{
	Use:   "api-resources",
	Short: "Print the supported API resources on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiResourcesCmd).Standalone()

	apiResourcesCmd.Flags().String("api-group", "", "Limit to resources in the specified API group.")
	apiResourcesCmd.Flags().Bool("cached", false, "Use the cached list of resources if available.")
	apiResourcesCmd.Flags().StringSlice("categories", nil, "Limit to resources that belong to the specified categories.")
	apiResourcesCmd.Flags().Bool("namespaced", false, "If false, non-namespaced resources will be returned, otherwise returning namespaced resources by default.")
	apiResourcesCmd.Flags().Bool("no-headers", false, "When using the default or custom-column output format, don't print headers (default print headers).")
	apiResourcesCmd.Flags().StringP("output", "o", "", "Output format. One of: (wide, name).")
	apiResourcesCmd.Flags().String("sort-by", "", "If non-empty, sort list of resources using specified field. The field can be either 'name' or 'kind'.")
	apiResourcesCmd.Flags().StringSlice("verbs", nil, "Limit to resources that support the specified verbs.")
	rootCmd.AddCommand(apiResourcesCmd)

	carapace.Gen(apiResourcesCmd).FlagCompletion(carapace.ActionMap{
		"api-group": kubectl.ActionApiGroups(),
		"output":    carapace.ActionValues("wide", "name"),
		"sort-by":   carapace.ActionValues("name", "kind"),
		"verbs":     kubectl.ActionResourceVerbs(),
	})
}
