package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace matching substrings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replaceCmd).Standalone()

	replaceCmd.Flags().BoolP("all", "a", false, "replace all matches")
	replaceCmd.Flags().BoolP("filter", "f", false, "only print strings where replacement was made")
	replaceCmd.Flags().BoolP("ignore-case", "i", false, "case-insensitive matching")
	replaceCmd.Flags().StringP("max-matches", "m", "", "stop after MAX matching lines")
	replaceCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	replaceCmd.Flags().BoolP("regex", "r", false, "interpret pattern as PCRE2 regex")

	rootCmd.AddCommand(replaceCmd)
}
