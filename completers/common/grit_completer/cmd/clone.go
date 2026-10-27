package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Copy a remote repository into a new directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().StringP("branch", "b", "", "Branch or tag to check out instead of the remote default")
	cloneCmd.Flags().String("depth", "", "Truncate fetched history to this many commits (`git clone --depth`)")
	cloneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cloneCmd.Flags().Bool("shallow", false, "Shallow clone of the fetched ref (`git clone --depth=1`; hidden alias)")
	cloneCmd.Flag("shallow").Hidden = true
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		carapace.ActionDirectories(),
	)
}
