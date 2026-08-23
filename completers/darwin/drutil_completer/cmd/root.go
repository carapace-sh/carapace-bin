package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "drutil",
	Short: "interact with CD/DVD burners",
	Long:  "https://keith.github.io/xcode-manpages/drutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("drive", "", "Specify a drive")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("burn", "erase", "eject", "help", "info", "list", "status", "tray"),
	)
}
