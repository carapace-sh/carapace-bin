package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwpolicy",
	Short: "gets and sets password policies",
	Long:  "https://keith.github.io/xcode-manpages/pwpolicy.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("a", "a", "", "Authenticator name")
	rootCmd.Flags().StringS("c", "c", "", "Computer account name")
	rootCmd.Flags().BoolS("h", "h", false, "Print help")
	rootCmd.Flags().StringS("n", "n", "", "Nodename")
	rootCmd.Flags().StringS("p", "p", "", "Password")
	rootCmd.Flags().StringS("u", "u", "", "Username")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"u": os.ActionUsers(),
	})
}
