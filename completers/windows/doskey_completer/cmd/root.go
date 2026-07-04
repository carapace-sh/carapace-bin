package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doskey",
	Short: "edit command lines, recall commands, and create macros",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/doskey",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("history", false, "display all commands stored in memory")
	rootCmd.Flags().Bool("insert", false, "new text is typed in the middle of old text")
	rootCmd.Flags().Bool("listsize", false, "set maximum number of commands stored in memory")
	rootCmd.Flags().Bool("overstrike", false, "new text overwrites old text")
	rootCmd.Flags().Bool("reinstall", false, "reinstall doskey")
}
