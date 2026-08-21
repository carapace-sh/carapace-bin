package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var listNewestCmd = &cobra.Command{
	Use:     "list-newest <repo?>",
	Aliases: []string{"ln"},
	Short:   "list the newest packages in the repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listNewestCmd).Standalone()

	listNewestCmd.Flags().IntP("last", "l", 0, "only show the newest packages since the nth repository update")
	listNewestCmd.Flags().StringP("since", "s", "", "show the newest since the specified date (YYYY-MM-DD)")

	rootCmd.AddCommand(listNewestCmd)

	carapace.Gen(listNewestCmd).PositionalAnyCompletion(
		eopkg.ActionRepositories().FilterArgs(),
	)
}
