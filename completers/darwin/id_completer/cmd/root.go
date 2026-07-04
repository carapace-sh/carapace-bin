package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "id",
	Short: "return user identity",
	Long:  "https://keith.github.io/xcode-manpages/id.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Display process audit ID")
	rootCmd.Flags().BoolS("F", "F", false, "Display unique user ID")
	rootCmd.Flags().BoolS("G", "G", false, "Display group IDs")
	rootCmd.Flags().BoolS("P", "P", false, "Display process ID as password entry")
	rootCmd.Flags().BoolS("g", "g", false, "Display effective group ID")
	rootCmd.Flags().BoolS("n", "n", false, "Display name instead of ID")
	rootCmd.Flags().BoolS("p", "p", false, "Display login, gid, groups, uid")
	rootCmd.Flags().BoolS("r", "r", false, "Display real ID instead of effective")

	carapace.Gen(rootCmd).PositionalAnyCompletion(os.ActionUsers())
}
