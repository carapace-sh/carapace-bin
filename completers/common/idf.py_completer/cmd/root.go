package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "idf.py",
	Short: "ESP-IDF CLI build management tool",
	Long:  "https://github.com/espressif/esp-idf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("baud", "b", "", "Global baud rate for all idf.py subcommands")
	rootCmd.PersistentFlags().StringP("build-dir", "B", "", "Build directory")
	rootCmd.PersistentFlags().Bool("ccache", false, "Use ccache in build")
	rootCmd.PersistentFlags().BoolP("cmake-warn-uninitialized", "w", false, "Enable CMake uninitialized variable warnings")
	rootCmd.PersistentFlags().Bool("configdep", true, "Use configdep wrapper to optimize rebuild process")
	rootCmd.PersistentFlags().StringArrayP("define-cache-entry", "D", nil, "Create a cmake cache entry")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Only process arguments, but don't execute actions")
	rootCmd.PersistentFlags().StringP("generator", "G", "", "CMake generator")
	rootCmd.PersistentFlags().StringP("jobs", "j", "", "Number of parallel build jobs passed to the build tool")
	rootCmd.PersistentFlags().Bool("list-targets", false, "Print list of supported targets and exit")
	rootCmd.PersistentFlags().StringArray("local-storage-url", nil, "Mirror URLs")
	rootCmd.PersistentFlags().Bool("no-ccache", false, "Disable ccache in build")
	rootCmd.PersistentFlags().Bool("no-configdep", false, "Disable configdep wrapper")
	rootCmd.PersistentFlags().Bool("no-hints", false, "Disable hints on how to resolve errors and logging")
	rootCmd.PersistentFlags().BoolP("no-warnings", "n", false, "Disable CMake uninitialized variable warnings")
	rootCmd.PersistentFlags().StringP("port", "p", "", "Serial port")
	rootCmd.PersistentFlags().String("preset", "", "Configuration preset to use")
	rootCmd.PersistentFlags().Bool("preview", false, "Enable IDF features that are still in preview")
	rootCmd.PersistentFlags().String("profile", "", "Specifies the profile to use for this command")
	rootCmd.PersistentFlags().StringP("project-dir", "C", "", "Project directory")
	rootCmd.PersistentFlags().String("registry-url", "", "URL of the Component Registry")
	rootCmd.PersistentFlags().String("service-profile", "", "Specifies the profile to use for this command")
	rootCmd.PersistentFlags().StringArray("storage-url", nil, "Storage URLs")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose build output")
	rootCmd.PersistentFlags().Bool("version", false, "Show IDF version and exit")

	rootCmd.Flag("dry-run").Hidden = true
	rootCmd.Flag("service-profile").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"baud":        espidf.ActionBaudRates(),
		"build-dir":   carapace.ActionDirectories(),
		"generator":   carapace.ActionValues("Ninja", "Unix Makefiles"),
		"port":        espidf.ActionSerialPorts(),
		"preset":      espidf.ActionPresets(),
		"project-dir": carapace.ActionDirectories(),
	})
}
