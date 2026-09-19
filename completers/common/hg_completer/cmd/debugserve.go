package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugserveCmd = &cobra.Command{
	Use:    "debugserve",
	Short:  "run a server with advanced settings",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugserveCmd).Standalone()

	debugserveCmd.Flags().String("logiofd", "", "file descriptor to log server I/O to")
	debugserveCmd.Flags().String("logiofile", "", "file to log server I/O to")
	debugserveCmd.Flags().Bool("sshstdio", false, "run an SSH server bound to process handles")
	rootCmd.AddCommand(debugserveCmd)
}
