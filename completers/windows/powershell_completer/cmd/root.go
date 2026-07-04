package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Windows PowerShell 5.1 command-line interface",
	Long:  "https://learn.microsoft.com/en-us/powershell/scripting/developer/cmdlet-aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("command", "c", false, "execute a command string")
	rootCmd.Flags().BoolP("file", "f", false, "execute a script file")
	rootCmd.Flags().BoolP("noprofile", "noprofile", false, "do not load profile scripts")
	rootCmd.Flags().BoolP("nologo", "nologo", false, "do not display copyright banner")
	rootCmd.Flags().BoolP("noninteractive", "noninteractive", false, "non-interactive mode")
}
