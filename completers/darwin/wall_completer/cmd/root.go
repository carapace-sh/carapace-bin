package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wall",
	Short: "write a message to users",
	Long:  "https://keith.github.io/xcode-manpages/wall.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("g", "g", "", "Send messages to users in this group")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"g": os.ActionGroups(),
	})
}
