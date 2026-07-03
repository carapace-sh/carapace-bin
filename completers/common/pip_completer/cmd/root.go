package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pip",
	Short: "package manager for Python packages",
	Long:  "https://pip.pypa.io/en/stable/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("cache-dir", "", "Store the cache data in dir")
	rootCmd.PersistentFlags().String("cert", "", "Path to alternate CA bundle")
	rootCmd.PersistentFlags().String("client-cert", "", "Path to SSL client certificate, a single file")
	rootCmd.PersistentFlags().Bool("debug", false, "Let unhandled exceptions propagate outside the main subroutine")
	rootCmd.PersistentFlags().Bool("disable-pip-version-check", false, "")
	rootCmd.PersistentFlags().String("exists-action", "", "Default action when a path already exists")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help")
	rootCmd.PersistentFlags().Bool("isolated", false, "Run pip in an isolated mode")
	rootCmd.PersistentFlags().String("keyring-provider", "", "Enable the credential lookup via the keyring library")
	rootCmd.PersistentFlags().String("log", "", "Path to a verbose appending log")
	rootCmd.PersistentFlags().Bool("no-cache-dir", false, "Disable the cache")
	rootCmd.PersistentFlags().Bool("no-color", false, "Suppress colored output")
	rootCmd.PersistentFlags().Bool("no-input", false, "Disable prompting for input")
	rootCmd.PersistentFlags().Bool("no-python-version-warning", false, "")
	rootCmd.PersistentFlags().String("proxy", "", "Specify a proxy in the form")
	rootCmd.PersistentFlags().String("python", "", "Run pip with the specified Python interpreter")
	rootCmd.PersistentFlags().Bool("require-virtualenv", false, "Allow pip to only run in a virtual environment")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Give less output")
	rootCmd.PersistentFlags().String("retries", "", "Maximum number of retries each connection should")
	rootCmd.PersistentFlags().String("resume-retries", "", "Maximum attempts to resume or restart an incomplete download")
	rootCmd.PersistentFlags().String("timeout", "", "Set the socket timeout (default 15 seconds)")
	rootCmd.PersistentFlags().String("trusted-host", "", "Mark this host or host:port pair as trusted")
	rootCmd.PersistentFlags().String("use-deprecated", "", "Enable deprecated functionality")
	rootCmd.PersistentFlags().String("use-feature", "", "Enable new functionality")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Give more output")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Show version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"cert":      carapace.ActionFiles(),
		"exists-action": carapace.ActionValuesDescribed(
			"s", "switch",
			"i", "ignore",
			"w", "wipe",
			"b", "backup",
			"a", "abort",
		),
		"keyring-provider": carapace.ActionValues("auto", "disabled", "import", "subprocess"),
		"log":              carapace.ActionFiles(),
	})
}
