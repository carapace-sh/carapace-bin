package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "networkquality",
	Short: "network quality testing tool",
	Long:  "https://keith.github.io/xcode-manpages/networkQuality.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bonjour", "B", "", "Run against specified Bonjour instance")
	rootCmd.Flags().BoolP("bonjour-list", "b", false, "Show Bonjour networkQuality servers")
	rootCmd.Flags().StringP("configuration", "C", "", "Override Configuration URL or path")
	rootCmd.Flags().StringP("force-protocols", "f", "", "Enforce protocol selections")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().BoolP("insecure", "k", false, "Disable certificate validation")
	rootCmd.Flags().StringP("interface", "I", "", "Bind test to interface")
	rootCmd.Flags().StringP("max-runtime", "M", "", "Provide maximum runtime in seconds")
	rootCmd.Flags().BoolP("no-download", "d", false, "Do not run a download test")
	rootCmd.Flags().BoolP("no-upload", "u", false, "Do not run an upload test")
	rootCmd.Flags().StringP("output", "c", "", "Produce computer-readable output to filename")
	rootCmd.Flags().StringP("override-host", "r", "", "Connect to host or IP, overriding DNS")
	rootCmd.Flags().BoolP("private-relay", "p", false, "Use iCloud Private Relay")
	rootCmd.Flags().BoolP("sequential", "s", false, "Run tests sequentially")
	rootCmd.Flags().StringP("server-port", "S", "", "Start and run server on specified port")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"force-protocols": carapace.ActionValues("h1", "h2", "h3", "L4S", "noL4S"),
		"output":          carapace.ActionFiles(),
	})
}