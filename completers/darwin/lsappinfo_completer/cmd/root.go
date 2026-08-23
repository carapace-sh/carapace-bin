package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsappinfo",
	Short: "list application information",
	Long:  "https://man.freebsd.org/cgi/man.cgi?lsappinfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(frontCmd)
	rootCmd.AddCommand(findCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listenCmd)
	rootCmd.AddCommand(launchCmd)
	rootCmd.AddCommand(metainfoCmd)
	rootCmd.AddCommand(processListCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(sharedmemoryCmd)
	rootCmd.AddCommand(unlistenCmd)
	rootCmd.AddCommand(visibleProcessListCmd)
	rootCmd.AddCommand(allocateASNCmd)
	rootCmd.AddCommand(createFileCmd)
	rootCmd.AddCommand(disconnectCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(foreverCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(removeFileCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(setinfoCmd)
	rootCmd.AddCommand(setmetainfoCmd)
	rootCmd.AddCommand(waitCmd)
	rootCmd.AddCommand(writePIDToFileCmd)
}
