package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Match strings against a pattern",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(matchCmd).Standalone()

	matchCmd.Flags().BoolP("all", "a", false, "report all matches")
	matchCmd.Flags().BoolP("entire", "e", false, "print entire matching string")
	matchCmd.Flags().BoolP("groups-only", "g", false, "report only capturing groups")
	matchCmd.Flags().BoolP("ignore-case", "i", false, "case-insensitive matching")
	matchCmd.Flags().BoolP("index", "n", false, "report 1-based start position and length")
	matchCmd.Flags().BoolP("invert", "v", false, "select lines that do not match")
	matchCmd.Flags().StringP("max-matches", "m", "", "stop after MAX matching lines")
	matchCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	matchCmd.Flags().BoolP("regex", "r", false, "interpret pattern as PCRE2 regex")

	rootCmd.AddCommand(matchCmd)
}
