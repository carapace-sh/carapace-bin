package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "start a new instance of the Windows command interpreter",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cmd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("c", "c", false, "carry out command and terminate")
	rootCmd.Flags().BoolP("k", "k", false, "carry out command and remain")
	rootCmd.Flags().BoolP("q", "q", false, "disable echo")
}
