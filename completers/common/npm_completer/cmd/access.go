package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Set access level on published packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessCmd).Standalone()
	accessCmd.PersistentFlags().Bool("json", false, "output as json")
	accessCmd.PersistentFlags().String("otp", "", "one-time password")
	accessCmd.PersistentFlags().String("registry", "", "base URL of the npm registry")

	rootCmd.AddCommand(accessCmd)

	carapace.Gen(accessCmd).FlagCompletion(carapace.ActionMap{
		"registry": carapace.ActionValues(),
	})
}
