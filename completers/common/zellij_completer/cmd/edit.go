package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit file with default $EDITOR / $VISUAL Returns: Created pane ID (format: terminal_<id>)",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().StringP("borderless", "b", "", "start this pane without a border (warning: will make it impossible to move with the mouse)")
	editCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	editCmd.Flags().String("cwd", "", "Change the working directory of the editor")
	editCmd.Flags().StringP("direction", "d", "", "Direction to open the new pane in")
	editCmd.Flags().BoolP("floating", "f", false, "Open the new pane in floating mode")
	editCmd.Flags().String("height", "", "The height if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	editCmd.Flags().BoolP("help", "h", false, "Print help")
	editCmd.Flags().BoolP("in-place", "i", false, "Open the new pane in place of the current pane, temporarily suspending it")
	editCmd.Flags().StringP("line-number", "l", "", "Open the file in the specified line number")
	editCmd.Flags().Bool("near-current-pane", false, "if set, will open the pane near the current one rather than following the user's focus")
	editCmd.Flags().Bool("no-focus", false, "if set, will open the pane without changing the focus of any client, placing it relative to the pane the command was issued from")
	editCmd.Flags().String("pinned", "", "Whether to pin a floating pane so that it is always on top")
	editCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	editCmd.Flags().String("width", "", "The width if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	editCmd.Flags().StringP("x", "x", "", "The x coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	editCmd.Flags().StringP("y", "y", "", "The y coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{
		"borderless": carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"cwd":        carapace.ActionDirectories(),
		"direction":  zellij.ActionDirections(),
		"pinned":     carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"tab-id":     zellij.ActionTabs(),
	})

	carapace.Gen(editCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
