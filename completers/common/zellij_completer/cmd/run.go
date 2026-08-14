package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run a command in a new pane Returns: Created pane ID (format: terminal_<id>)",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("block-until-exit", false, "Block until the command exits (regardless of exit status) OR its pane has been closed")
	runCmd.Flags().Bool("block-until-exit-failure", false, "Block until the command exits with failure (non-zero exit status) OR its pane has been closed")
	runCmd.Flags().Bool("block-until-exit-success", false, "Block until the command exits successfully (exit status 0) OR its pane has been closed")
	runCmd.Flags().Bool("blocking", false, "Block until the command has finished and its pane has been closed")
	runCmd.Flags().StringP("borderless", "b", "", "start this pane without a border (warning: will make it impossible to move with the mouse)")
	runCmd.Flags().BoolP("close-on-exit", "c", false, "Close the pane immediately when its command exits")
	runCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	runCmd.Flags().String("cwd", "", "Change the working directory of the new pane")
	runCmd.Flags().StringP("direction", "d", "", "Direction to open the new pane in")
	runCmd.Flags().BoolP("floating", "f", false, "Open the new pane in floating mode")
	runCmd.Flags().String("height", "", "The height if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	runCmd.Flags().BoolP("help", "h", false, "Print help")
	runCmd.Flags().BoolP("in-place", "i", false, "Open the new pane in place of the current pane, temporarily suspending it")
	runCmd.Flags().StringP("name", "n", "", "Name of the new pane")
	runCmd.Flags().Bool("near-current-pane", false, "if set, will open the pane near the current one rather than following the user's focus")
	runCmd.Flags().Bool("no-focus", false, "if set, will open the pane without changing the focus of any client, placing it relative to the pane the command was issued from")
	runCmd.Flags().String("pinned", "", "Whether to pin a floating pane so that it is always on top")
	runCmd.Flags().Bool("stacked", false, "")
	runCmd.Flags().BoolP("start-suspended", "s", false, "Start the command suspended, only running after you first presses ENTER")
	runCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	runCmd.Flags().String("width", "", "The width if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	runCmd.Flags().StringP("x", "x", "", "The x coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	runCmd.Flags().StringP("y", "y", "", "The y coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"borderless": carapace.ActionValues("true", "false"),
		"cwd":        carapace.ActionFiles(),
		"direction":  actionDirections(),
		"pinned":     carapace.ActionValues("true", "false"),
	})
}
