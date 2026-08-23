package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crontab",
	Short: "maintain crontab files for individual users",
	Long:  "https://keith.github.io/xcode-manpages/crontab.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("e", "e", false, "Edit the current crontab")
	rootCmd.Flags().BoolS("l", "l", false, "Display the current crontab")
	rootCmd.Flags().BoolS("r", "r", false, "Remove the current crontab")
	rootCmd.Flags().StringS("u", "u", "", "Specify the name of the user whose crontab is to be tweaked")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"u": os.ActionUsers(),
	})
}
