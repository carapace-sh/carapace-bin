package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasklist",
	Short: "display a list of currently running processes",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tasklist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("m", "m", false, "show module information for each process")
	rootCmd.Flags().BoolP("svc", "svc", false, "display services hosted in each process")
	rootCmd.Flags().BoolP("v", "v", false, "display verbose task information")
}
