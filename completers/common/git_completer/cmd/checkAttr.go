package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var checkAttrCmd = &cobra.Command{
	Use:     "check-attr",
	Short:   "Display gitattributes information",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(checkAttrCmd).Standalone()

	checkAttrCmd.Flags().BoolP("all", "a", false, "report all attributes set on file")
	checkAttrCmd.Flags().Bool("cached", false, "use .gitattributes only from the index")
	checkAttrCmd.Flags().String("source", "", "which tree-ish to check attributes at")
	checkAttrCmd.Flags().Bool("stdin", false, "read file names from stdin")
	checkAttrCmd.Flags().BoolS("z", "z", false, "terminate input and output records by a NUL character")
	rootCmd.AddCommand(checkAttrCmd)

	carapace.Gen(checkAttrCmd).FlagCompletion(carapace.ActionMap{
		"source": git.ActionRefs(git.RefOption{}.Default()),
	})

	// TODO complete available attributes of ref

	carapace.Gen(checkAttrCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := checkAttrCmd.Flag("source"); f.Changed {
				return git.ActionRefFiles(f.Value.String())
			}
			return carapace.ActionFiles()
		}),
	)
}
