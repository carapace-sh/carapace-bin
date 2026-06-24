package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nameRevCmd = &cobra.Command{
	Use:     "name-rev",
	Short:   "Find symbolic names for given revs",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(nameRevCmd).Standalone()

	nameRevCmd.Flags().Bool("all", false, "list all commits reachable from all refs")
	nameRevCmd.Flags().Bool("always", false, "show uniquely abbreviated commit object as fallback")
	nameRevCmd.Flags().Bool("annotate-stdin", false, "transform stdin by substituting all the 40-character SHA-1 hexes")
	nameRevCmd.Flags().StringArray("exclude", nil, "do not use any ref whose name matches a given shell pattern")
	nameRevCmd.Flags().Bool("name-only", false, "print only the name")
	nameRevCmd.Flags().StringArray("refs", nil, "only use refs whose names match a given shell pattern")
	nameRevCmd.Flags().Bool("tags", false, "only use tags to name the commits")
	nameRevCmd.Flags().Bool("undefined", false, "allow to print undefined names")
	rootCmd.AddCommand(nameRevCmd)
}
