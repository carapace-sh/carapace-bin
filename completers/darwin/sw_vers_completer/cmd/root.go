package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sw_vers",
	Short: "print operating system version information",
	Long:  "https://keith.github.io/xcode-manpages/sw_vers.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("buildVersion", "buildVersion", "", "Print the build version")
	rootCmd.Flags().StringP("productName", "productName", "", "Print the product name")
	rootCmd.Flags().StringP("productVersion", "productVersion", "", "Print the product version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"buildVersion":   carapace.ActionValues().StyleF(style.ForKeyword),
		"productName":    carapace.ActionValues().StyleF(style.ForKeyword),
		"productVersion": carapace.ActionValues().StyleF(style.ForKeyword),
	})
}
