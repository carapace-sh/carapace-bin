package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var listAvailableCmd = &cobra.Command{
	Use:     "list-available <repo-name?>",
	Aliases: []string{"la"},
	Short:   "list all available packages in all repositories, or just in the repositories specified",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listAvailableCmd).Standalone()

	listAvailableCmd.Flags().StringP("component", "c", "", "list available packages under the given component")
	listAvailableCmd.Flags().BoolP("long", "l", false, "use long output instead of brief one line descriptions")
	listAvailableCmd.Flags().BoolP("uninstalled", "U", false, "only show uninstalled packages")

	rootCmd.AddCommand(listAvailableCmd)

	carapace.Gen(listAvailableCmd).PositionalAnyCompletion(
		eopkg.ActionRepositories().FilterArgs(),
	)
}
