package cmd

import (
	"github.com/carapace-sh/carapace"
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

	rootCmd.Flags().BoolP("buildVersion", "buildVersion", false, "Print the build version")
	rootCmd.Flags().BoolP("productName", "productName", false, "Print the product name")
	rootCmd.Flags().BoolP("productVersion", "productVersion", false, "Print the product version")

}
