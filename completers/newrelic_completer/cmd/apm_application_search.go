package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apm_application_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a New Relic application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_application_searchCmd).Standalone()
	apm_application_searchCmd.Flags().StringP("name", "n", "", "search for results matching the given APM application name")
	apm_applicationCmd.AddCommand(apm_application_searchCmd)
}
