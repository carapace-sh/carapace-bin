package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a shell command in the context of a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	execCmd.Flags().BoolP("shell-mode", "c", false, "Run the command inside of a shell. Uses `/bin/sh` on UNIX and `cmd.exe` on Windows")
	rootCmd.AddCommand(execCmd)
}
