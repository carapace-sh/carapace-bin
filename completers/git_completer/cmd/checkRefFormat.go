package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkRefFormatCmd = &cobra.Command{
	Use:     "check-ref-format",
	Short:   "Ensures that a reference name is well formed",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(checkRefFormatCmd).Standalone()

	checkRefFormatCmd.Flags().Bool("allow-onelevel", false, "Controls whether one-level refnames are accepted")
	checkRefFormatCmd.Flags().Bool("branch", false, "Check for valid branch name")
	checkRefFormatCmd.Flags().Bool("no-allow-onelevel", false, "Controls whether one-level refnames are accepted")
	checkRefFormatCmd.Flags().Bool("normalize", false, "Normalize refname by removing any leading slash characters")
	checkRefFormatCmd.Flags().Bool("refspec-pattern", false, "Interpret <refname> as a reference name pattern for a refspec")
	rootCmd.AddCommand(checkRefFormatCmd)
}
