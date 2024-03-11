package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var columnCmd = &cobra.Command{
	Use:     "column",
	Short:   "Display data in columns",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(columnCmd).Standalone()

	columnCmd.Flags().String("command", "", "lookup config vars")
	columnCmd.Flags().String("indent", "", "padding space on left border")
	columnCmd.Flags().String("mode", "", "layout to use")
	columnCmd.Flags().String("nl", "", "padding space on right border")
	columnCmd.Flags().String("padding", "", "padding space between columns")
	columnCmd.Flags().String("raw-mode", "", "layout to use")
	columnCmd.Flags().String("width", "", "maximum width")
	rootCmd.AddCommand(columnCmd)

	carapace.Gen(columnCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.ActionCommands(rootCmd),
		"mode":    git.ActionColumnLayoutModes().UniqueList(","),
	})
}
