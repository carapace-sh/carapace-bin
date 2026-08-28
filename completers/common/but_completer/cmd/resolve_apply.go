package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolve_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Resolve conflicts of a conflicted commit, without entering resolution mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolve_applyCmd).Standalone()

	resolve_applyCmd.Flags().Bool("ai", false, "Let the configured AI model merge the targeted conflicts")
	resolve_applyCmd.Flags().String("commit", "", "A conflicted commit, or a branch (meaning its oldest conflicted commit — branch names stay stable across applies, unlike commit ids). Defaults to the first conflicted branch's oldest conflicted commit")
	resolve_applyCmd.Flags().StringP("file", "F", "", "Read the replacement content from this file (otherwise from stdin)")
	resolve_applyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolve_applyCmd.Flags().Bool("ours", false, "Take the ours side: the new base the commit was rebased onto")
	resolve_applyCmd.Flags().Bool("theirs", false, "Take the theirs side: the commit's own version")
	resolveCmd.AddCommand(resolve_applyCmd)
}