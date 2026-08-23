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

	rootCmd.Flags().Bool("buildVersion", false, "Print the build version")
	rootCmd.Flags().Bool("productName", false, "Print the product name")
	rootCmd.Flags().Bool("productVersion", false, "Print the product version")
	rootCmd.Flags().Bool("productVersionExtra", false, "Print the product version extra")

}
