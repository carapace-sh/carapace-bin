package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

func addSizeFlags(cmd *cobra.Command) {
	cmd.Flags().String("diff", "", "Show the differences in comparison with another project")
	cmd.Flags().String("format", "", "Specify output format")
	cmd.Flags().BoolP("legacy", "l", false, "Use legacy esp-idf-size version")
	cmd.Flags().String("output-file", "", "Print output to the specified file instead of stdout")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"diff":        carapace.ActionFiles(),
		"format":      carapace.ActionValues("csv", "default", "json", "json2", "raw", "text", "tree"),
		"output-file": carapace.ActionFiles(),
	})
}

func init() {
	allCmd := &cobra.Command{
		Use:     "all",
		Aliases: []string{"build"},
		Short:   "Build the project",
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(allCmd).Standalone()
	rootCmd.AddCommand(allCmd)

	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Build only the app",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(appCmd).Standalone()
	rootCmd.AddCommand(appCmd)

	bootloaderCmd := &cobra.Command{
		Use:   "bootloader",
		Short: "Build only bootloader",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(bootloaderCmd).Standalone()
	rootCmd.AddCommand(bootloaderCmd)

	buildSystemTargetsCmd := &cobra.Command{
		Use:   "build-system-targets",
		Short: "Print list of build system targets",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(buildSystemTargetsCmd).Standalone()
	rootCmd.AddCommand(buildSystemTargetsCmd)

	cleanCmd := &cobra.Command{
		Use:   "clean",
		Short: "Delete build output files from the build directory",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cleanCmd).Standalone()
	rootCmd.AddCommand(cleanCmd)

	configReportCmd := &cobra.Command{
		Use:   "config-report",
		Short: "Generate a JSON file with the project configuration report",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(configReportCmd).Standalone()
	rootCmd.AddCommand(configReportCmd)

	confserverCmd := &cobra.Command{
		Use:   "confserver",
		Short: "Run JSON configuration server",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(confserverCmd).Standalone()
	confserverCmd.Flags().String("buffer-size", "", "Set the buffer size (in KB)")
	rootCmd.AddCommand(confserverCmd)

	docsCmd := &cobra.Command{
		Use:   "docs",
		Short: "Open web browser with documentation for ESP-IDF",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(docsCmd).Standalone()
	docsCmd.Flags().StringP("language", "l", "", "Documentation language")
	docsCmd.Flags().Bool("no-browser", false, "Don't open browser")
	docsCmd.Flags().String("starting-page", "", "Documentation page")
	docsCmd.Flags().StringP("target", "t", "", "Chip target")
	docsCmd.Flags().String("version", "", "Version of ESP-IDF")
	carapace.Gen(docsCmd).FlagCompletion(carapace.ActionMap{
		"language": carapace.ActionValues("cn", "en", "zh_CN"),
		"target":   espidf.ActionTargets(),
	})
	rootCmd.AddCommand(docsCmd)

	efuseCommonTableCmd := &cobra.Command{
		Use:   "efuse-common-table",
		Short: "Generate C-source for IDF's eFuse fields",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseCommonTableCmd).Standalone()
	rootCmd.AddCommand(efuseCommonTableCmd)

	efuseCustomTableCmd := &cobra.Command{
		Use:   "efuse-custom-table",
		Short: "Generate C-source for user's eFuse fields",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseCustomTableCmd).Standalone()
	rootCmd.AddCommand(efuseCustomTableCmd)

	fullcleanCmd := &cobra.Command{
		Use:   "fullclean",
		Short: "Delete the entire build directory contents",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(fullcleanCmd).Standalone()
	rootCmd.AddCommand(fullcleanCmd)

	menuconfigCmd := &cobra.Command{
		Use:   "menuconfig",
		Short: "Run menuconfig project configuration tool",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(menuconfigCmd).Standalone()
	menuconfigCmd.Flags().String("color-scheme", "", "Menuconfig color scheme")
	menuconfigCmd.Flags().String("style", "", "Menuconfig color scheme")
	carapace.Gen(menuconfigCmd).FlagCompletion(carapace.ActionMap{
		"color-scheme": carapace.ActionValues("dark", "light"),
		"style":        carapace.ActionValues("dark", "light"),
	})
	rootCmd.AddCommand(menuconfigCmd)

	partitionTableCmd := &cobra.Command{
		Use:   "partition-table",
		Short: "Build only partition table",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(partitionTableCmd).Standalone()
	rootCmd.AddCommand(partitionTableCmd)

	pythonCleanCmd := &cobra.Command{
		Use:   "python-clean",
		Short: "Delete generated Python byte code from the IDF directory",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(pythonCleanCmd).Standalone()
	rootCmd.AddCommand(pythonCleanCmd)

	reconfigureCmd := &cobra.Command{
		Use:   "reconfigure",
		Short: "Re-run CMake",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(reconfigureCmd).Standalone()
	rootCmd.AddCommand(reconfigureCmd)

	refreshConfigCmd := &cobra.Command{
		Use:   "refresh-config",
		Short: "Resolve conflicts in default values of config options",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(refreshConfigCmd).Standalone()
	refreshConfigCmd.Flags().String("policy", "", "Policy for handling defaults in the configuration")
	carapace.Gen(refreshConfigCmd).FlagCompletion(carapace.ActionMap{
		"policy": carapace.ActionValues("interactive", "kconfig", "sdkconfig"),
	})
	rootCmd.AddCommand(refreshConfigCmd)

	saveDefconfigCmd := &cobra.Command{
		Use:   "save-defconfig",
		Short: "Generate a sdkconfig.defaults with minimal options",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(saveDefconfigCmd).Standalone()
	saveDefconfigCmd.Flags().Bool("add-menu-labels", false, "Add menu labels to minimal config")
	rootCmd.AddCommand(saveDefconfigCmd)

	setTargetCmd := &cobra.Command{
		Use:   "set-target",
		Short: "Set the chip target to build",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(setTargetCmd).Standalone()
	carapace.Gen(setTargetCmd).PositionalCompletion(
		espidf.ActionTargets(),
	)
	rootCmd.AddCommand(setTargetCmd)

	showEfuseTableCmd := &cobra.Command{
		Use:   "show-efuse-table",
		Short: "Print eFuse table",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(showEfuseTableCmd).Standalone()
	rootCmd.AddCommand(showEfuseTableCmd)

	sizeCmd := &cobra.Command{
		Use:   "size",
		Short: "Print basic size information about the app",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(sizeCmd).Standalone()
	addSizeFlags(sizeCmd)
	rootCmd.AddCommand(sizeCmd)

	sizeComponentsCmd := &cobra.Command{
		Use:   "size-components",
		Short: "Print per-component size information",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(sizeComponentsCmd).Standalone()
	addSizeFlags(sizeComponentsCmd)
	rootCmd.AddCommand(sizeComponentsCmd)

	sizeFilesCmd := &cobra.Command{
		Use:   "size-files",
		Short: "Print per-source-file size information",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(sizeFilesCmd).Standalone()
	addSizeFlags(sizeFilesCmd)
	rootCmd.AddCommand(sizeFilesCmd)
}
