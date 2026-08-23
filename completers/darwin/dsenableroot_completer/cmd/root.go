package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dsenableroot",
	Short: "enables or disables the root account",
	Long:  "https://keith.github.io/xcode-manpages/dsenableroot.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "Disable the root account")
	rootCmd.Flags().StringS("p", "p", "", "Password")
	rootCmd.Flags().StringS("r", "r", "", "Root password")
	rootCmd.Flags().StringS("u", "u", "", "Username")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"u": os.ActionUsers(),
	})
}
