package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/softwareupdate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "softwareupdate",
	Short: "system software update tool",
	Long:  "https://keith.github.io/xcode-manpages/softwareupdate.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("agree-to-license", false, "Agree to the software license agreement without user interaction")
	rootCmd.Flags().BoolP("all", "a", false, "Install all available updates")
	rootCmd.Flags().Bool("background", false, "Trigger a background scan and update operation")
	rootCmd.Flags().BoolP("download", "d", false, "Download updates without installing")
	rootCmd.Flags().Bool("dump-state", false, "Log the internal state of the SU daemon")
	rootCmd.Flags().Bool("evaluate-products", false, "Evaluate a list of product keys")
	rootCmd.Flags().Bool("fetch-full-installer", false, "Fetch the full macOS installer")
	rootCmd.Flags().String("force", "", "Force an operation to complete")
	rootCmd.Flags().String("full-installer-version", "", "The version of macOS to install")
	rootCmd.Flags().Bool("help", false, "Display usage information")
	rootCmd.Flags().Bool("history", false, "Show the install history")
	rootCmd.Flags().BoolP("install", "i", false, "Install updates")
	rootCmd.Flags().Bool("install-rosetta", false, "Install Rosetta 2")
	rootCmd.Flags().BoolP("list", "l", false, "List all appropriate update labels")
	rootCmd.Flags().Bool("list-full-installers", false, "List all available full installers")
	rootCmd.Flags().Bool("no-scan", false, "Do not scan when listing or installing updates")
	rootCmd.Flags().Bool("os-only", false, "Only install OS updates")
	rootCmd.Flags().String("product-types", "", "Limit a scan to a particular product type")
	rootCmd.Flags().String("products", "", "A comma-separated list of product keys to operate on")
	rootCmd.Flags().BoolP("recommended", "r", false, "Install recommended updates only")
	rootCmd.Flags().BoolP("restart", "R", false, "Automatically restart or shut down if required")
	rootCmd.Flags().Bool("safari-only", false, "Only install Safari updates")
	rootCmd.Flags().Bool("schedule", false, "Set or check the scheduled update setting")
	rootCmd.Flags().Bool("stdinpass", false, "Read password from stdin")
	rootCmd.Flags().String("user", "", "Local username to authenticate as an owner")
	rootCmd.Flags().Bool("verbose", false, "Enable verbose output")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		softwareupdate.ActionUpdates(),
	)
}
