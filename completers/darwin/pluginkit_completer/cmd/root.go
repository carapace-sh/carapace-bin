package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pluginkit",
	Short: "plugin plug-in extension pluginkit",
	Long:  "https://man.freebsd.org/cgi/man.cgi?pluginkit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("a", false, "Add plugins at file location")
	rootCmd.Flags().Bool("all-versions", false, "Find all versions of a given plug-in")
	rootCmd.Flags().Bool("duplicates", false, "Find all physical instances")
	rootCmd.Flags().String("election", "", "Apply user election setting")
	rootCmd.Flags().String("identifier", "", "Plug-in identifier to match")
	rootCmd.Flags().Bool("match", false, "Scan all registered plug-ins")
	rootCmd.Flags().String("platform", "", "Plug-in platform")
	rootCmd.Flags().String("protocol", "", "Plug-in protocol to match")
	rootCmd.Flags().Bool("r", false, "Remove plugins at file location")
	rootCmd.Flags().Bool("verbose", false, "Verbose")
	rootCmd.Flags().Bool("version", false, "Version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"election": carapace.ActionValues("use", "ignore", "default"),
		"platform": carapace.ActionValues("native", "maccatalyst"),
	})
}
