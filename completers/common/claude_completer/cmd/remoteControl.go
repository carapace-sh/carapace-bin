package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remoteControlCmd = &cobra.Command{
	Use:   "remote-control",
	Short: "Start a Remote Control server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteControlCmd).Standalone()

	remoteControlCmd.Flags().String("capacity", "", "Maximum concurrent sessions")
	remoteControlCmd.Flags().Bool("chrome", false, "Enable Claude in Chrome for spawned sessions")
	remoteControlCmd.Flags().BoolP("continue", "c", false, "Reattach to the last recorded session for this directory")
	remoteControlCmd.Flags().Bool("create-session-in-dir", false, "Pre-create a session in the current directory")
	remoteControlCmd.Flags().StringP("debug", "d", "", "Enable debug mode with optional category filtering")
	remoteControlCmd.Flags().String("debug-file", "", "Write debug logs to a file")
	remoteControlCmd.Flags().BoolP("help", "h", false, "Display help for command")
	remoteControlCmd.Flags().String("name", "", "Name for the session")
	remoteControlCmd.Flags().Bool("no-chrome", false, "Disable Claude in Chrome for spawned sessions")
	remoteControlCmd.Flags().Bool("no-create-session-in-dir", false, "Do not pre-create a session in the current directory")
	remoteControlCmd.Flags().String("permission-mode", "", "Permission mode for spawned sessions")
	remoteControlCmd.Flags().String("remote-control-session-name-prefix", "", "Prefix for auto-generated session names")
	remoteControlCmd.Flags().String("session-id", "", "Reattach to a specific session by ID")
	remoteControlCmd.Flags().String("spawn", "", "Spawn mode for on-demand sessions")
	remoteControlCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	remoteControlCmd.Flag("debug").NoOptDefVal = " "
	rootCmd.AddCommand(remoteControlCmd)

	carapace.Gen(remoteControlCmd).FlagCompletion(carapace.ActionMap{
		"debug-file":      carapace.ActionFiles(),
		"permission-mode": carapace.ActionValues("acceptEdits", "auto", "bypassPermissions", "default", "dontAsk", "plan"),
		"spawn":           carapace.ActionValues("same-dir", "worktree", "session"),
	})
}
