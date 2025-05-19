package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rangeDiffCmd = &cobra.Command{
	Use:     "range-diff",
	Short:   "Compare two commit ranges",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(rangeDiffCmd).Standalone()

	rangeDiffCmd.Flags().String("creation-factor", "", "Set the creation/deletion cost fudge factor to <percent>")
	rangeDiffCmd.Flags().Bool("no-dual-color", false, "Do not recreate the original diffs’ coloring")
	rangeDiffCmd.Flags().String("no-notes", "", "This flag is passed to the git log progeram")
	rangeDiffCmd.Flags().String("notes", "", "This flag is passed to the git log program")
	common.AddDiffFlags(rangeDiffCmd)
	rootCmd.AddCommand(rangeDiffCmd)

	rangeDiffCmd.Flag("no-notes").NoOptDefVal = " "
	rangeDiffCmd.Flag("notes").NoOptDefVal = " "

	carapace.Gen(rangeDiffCmd).FlagCompletion(carapace.ActionMap{
		"no-notes": git.ActionRefs(git.RefOption{}.Default()),
		"notes":    git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(rangeDiffCmd).PositionalCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.Contains(c.Args[0], "...") {
				return git.ActionRefs(git.RefOption{}.Default())
			}
			return carapace.ActionValues()
		}),
	)
}
