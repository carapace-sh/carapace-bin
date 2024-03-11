package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenses_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Check the licenses of the installed packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenses_listCmd).Standalone()

	licenses_listCmd.Flags().BoolP("dev", "D", false, "Check only \"devDependencies\"")
	licenses_listCmd.Flags().Bool("json", false, "Show information in JSON format")
	licenses_listCmd.Flags().Bool("long", false, "Show more details (such as a link to the repo) are")
	licenses_listCmd.Flags().Bool("no-optional", false, "Don't check \"optionalDependencies\"")
	licenses_listCmd.Flags().BoolP("prod", "P", false, "Check only \"dependencies\" and \"optionalDependencies\"")
	licensesCmd.AddCommand(licenses_listCmd)
}
