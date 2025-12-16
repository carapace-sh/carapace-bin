package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-coauthor",
	Short: "Add a co-author to the last commit",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-coauthor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalCompletion(
		gh.ActionUsers(gh.HostOpts{}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionUserEmails(gh.OwnerOpts{Owner: c.Args[0]})
		}),
	)
}
