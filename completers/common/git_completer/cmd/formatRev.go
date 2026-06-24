package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var formatRevCmd = &cobra.Command{
	Use:     "format-rev",
	Short:   "EXPERIMENTAL: Pretty format revisions on demand",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(formatRevCmd).Standalone()

	formatRevCmd.Flags().String("format", "", "Pretty format to use")
	formatRevCmd.Flags().StringArray("notes", nil, "Display notes for pretty format")
	formatRevCmd.Flags().BoolP("null", "z", false, "Use NUL for input and output termination")
	formatRevCmd.Flags().Bool("null-input", false, "Use NUL for input termination")
	formatRevCmd.Flags().Bool("null-output", false, "Use NUL for output termination")
	formatRevCmd.Flags().String("stdin-mode", "", "How revs are processed")
	rootCmd.AddCommand(formatRevCmd)

	formatRevCmd.Flag("notes").NoOptDefVal = " "

	carapace.Gen(formatRevCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"notes":      git.ActionRefs(git.RefOption{}.Default()),
		"stdin-mode": carapace.ActionValues("walk", "single"),
	})
}
