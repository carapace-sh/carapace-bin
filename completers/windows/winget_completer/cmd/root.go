package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "winget",
	Short: "Windows Package Manager",
	Long:  "https://github.com/microsoft/winget-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("disable-interactivity", false, "Disable interactive prompts")
	rootCmd.PersistentFlags().BoolP("help", "?", false, "Shows help about the selected command")
	rootCmd.PersistentFlags().Bool("ignore-warnings", false, "Suppresses warning outputs")
	rootCmd.Flags().Bool("info", false, "Display general info of the tool")
	rootCmd.PersistentFlags().Bool("logs", false, "Open the default logs location")
	rootCmd.PersistentFlags().Bool("no-proxy", false, "Disable the use of proxy for this execution")
	rootCmd.PersistentFlags().Bool("nowarn", false, "Suppresses warning outputs")
	rootCmd.PersistentFlags().Bool("open-logs", false, "Open the default logs location")
	rootCmd.PersistentFlags().String("proxy", "", "Set a proxy to use for this execution")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enables verbose logging for winget")
	rootCmd.PersistentFlags().Bool("verbose-logs", false, "Enables verbose logging for winget")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of the tool")
	rootCmd.PersistentFlags().Bool("wait", false, "Prompts the user to press any key before exiting")
}
