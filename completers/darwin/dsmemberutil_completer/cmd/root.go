package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dsmemberutil",
	Short: "membership API operations",
	Long:  "https://keith.github.io/xcode-manpages/dsmemberutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("checkmembership", "", "Check membership of user in group")
	rootCmd.Flags().String("getSID", "", "Get SID for user or group")
	rootCmd.Flags().String("getUUID", "", "Get UUID for user or group")
	rootCmd.Flags().BoolS("h", "h", false, "Print help")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"checkmembership": os.ActionUsers(),
	})
}
