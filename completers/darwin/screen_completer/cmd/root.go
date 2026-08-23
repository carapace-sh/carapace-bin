package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "screen",
	Short: "multi-screen window manager",
	Long:  "https://man.freebsd.org/cgi/man.cgi?screen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "Turn on automatic output logging")
	rootCmd.Flags().BoolS("T", "T", false, "Set the terminal type")
	rootCmd.Flags().BoolS("X", "X", false, "Execute a command in a specified session")
	rootCmd.Flags().BoolS("a", "a", false, "Include all capabilities")
	rootCmd.Flags().BoolS("c", "c", false, "Set the configuration file")
	rootCmd.Flags().BoolS("d", "d", false, "Detach a running session")
	rootCmd.Flags().BoolS("D", "D", false, "Detach and logout")
	rootCmd.Flags().BoolS("e", "e", false, "Set the command character")
	rootCmd.Flags().BoolS("h", "h", false, "Set the history scrollback size")
	rootCmd.Flags().BoolS("ls", "ls", false, "List sessions")
	rootCmd.Flags().BoolS("list", "list", false, "List sessions")
	rootCmd.Flags().BoolS("m", "m", false, "Force creation of a new session")
	rootCmd.Flags().BoolS("O", "O", false, "Select optimal output")
	rootCmd.Flags().BoolS("q", "q", false, "Check if session is running")
	rootCmd.Flags().BoolS("r", "r", false, "Reattach a session")
	rootCmd.Flags().BoolS("R", "R", false, "Reattach if possible, otherwise start a new session")
	rootCmd.Flags().BoolS("s", "s", false, "Set the shell")
	rootCmd.Flags().BoolS("S", "S", false, "Name a session")
	rootCmd.Flags().BoolS("t", "t", false, "Set the title")
	rootCmd.Flags().BoolS("U", "U", false, "Use UTF-8 encoding")
	rootCmd.Flags().BoolS("v", "v", false, "Version")
	rootCmd.Flags().BoolS("wipe", "wipe", false, "Wipe sessions")
	rootCmd.Flags().BoolS("x", "x", false, "Attach to a not detached screen session")
}
