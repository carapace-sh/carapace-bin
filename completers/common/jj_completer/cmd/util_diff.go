package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var util_diffCmd = &cobra.Command{
	Use:   "diff [OPTIONS] <PATH1> <PATH2>",
	Short: "Compare two files on disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_diffCmd).Standalone()

	util_diffCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	util_diffCmd.Flags().String("context", "", "Number of lines of context to show")
	util_diffCmd.Flags().Bool("git", false, "Show a Git-format diff")
	util_diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	util_diffCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines")
	util_diffCmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in amount of whitespace when comparing lines")
	util_diffCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	util_diffCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	util_diffCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or deleted")
	util_diffCmd.Flags().String("tool", "", "Generate diff by external command")
	util_diffCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	util_diffCmd.Flag("name-only").Hidden = true
	util_diffCmd.Flag("stat").Hidden = true
	util_diffCmd.Flag("summary").Hidden = true
	util_diffCmd.Flag("types").Hidden = true
	utilCmd.AddCommand(util_diffCmd)

	carapace.Gen(util_diffCmd).FlagCompletion(carapace.ActionMap{
		"tool": bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(util_diffCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
