package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orgCmd = &cobra.Command{
	Use:     "org",
	Short:   "Manage orgs",
	Aliases: []string{"ogr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orgCmd).Standalone()
	orgCmd.PersistentFlags().String("otp", "", "one-time password")
	orgCmd.PersistentFlags().Bool("json", false, "output as json")
	orgCmd.PersistentFlags().BoolP("parseable", "p", false, "output as json")
	orgCmd.PersistentFlags().String("registry", "", "base URL of the npm registry")

	rootCmd.AddCommand(orgCmd)
}
