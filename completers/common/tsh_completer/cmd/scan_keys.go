package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scan_keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Scan the local machine for SSH private keys and report findings to Teleport.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scan_keysCmd).Standalone()

	scan_keysCmd.Flags().String("dirs", "", "Directories to scan. Defaults to /home/ on Linux, /Users/ on macOS, and C:\\Users\\ on Windows.")
	scan_keysCmd.Flags().String("skip-paths", "", "Paths to directories or files to skip. Supports for matching patterns.")
	scanCmd.AddCommand(scan_keysCmd)
}
