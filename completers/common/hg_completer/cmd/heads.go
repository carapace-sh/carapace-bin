package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var headsCmd = &cobra.Command{
	Use:     "heads",
	Short:   "show branch heads",
	GroupID: groups[group_change_navigation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(headsCmd).Standalone()

	headsCmd.Flags().BoolP("active", "a", false, "show active branchheads only (DEPRECATED)")
	headsCmd.Flags().BoolP("closed", "c", false, "show normal and closed branch heads")
	headsCmd.Flags().StringP("rev", "r", "", "show only heads which are descendants of STARTREV")
	headsCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	headsCmd.Flags().StringP("template", "T", "", "display with template")
	headsCmd.Flags().BoolP("topo", "t", false, "show topological heads only")
	rootCmd.AddCommand(headsCmd)

	carapace.Gen(headsCmd).FlagCompletion(carapace.ActionMap{
		"rev": hg.ActionRevisions(),
	})
}
