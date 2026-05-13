package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swiftly",
	Short: "A utility for installing and managing Swift toolchains",
	Long:  "https://www.swift.org/install/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Perform swiftly initialization into your user account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var installCmd = &cobra.Command{
	Use:   "install [toolchain]",
	Short: "Install a new toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listAvailableCmd = &cobra.Command{
	Use:   "list-available [toolchain]",
	Short: "List toolchains available for install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var useCmd = &cobra.Command{
	Use:   "use [toolchain]",
	Short: "Set the in-use or default toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall <toolchain>...",
	Short: "Remove an installed toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listCmd = &cobra.Command{
	Use:   "list [toolchain]",
	Short: "List installed toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var updateCmd = &cobra.Command{
	Use:   "update [toolchain]",
	Short: "Update an installed toolchain to a newer version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update the version of swiftly itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var runCmd = &cobra.Command{
	Use:   "run <command>...",
	Short: "Run a command while proxying to the selected toolchain commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var selfUninstallCmd = &cobra.Command{
	Use:   "self-uninstall",
	Short: "Uninstall swiftly itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link swiftly so it resumes management of the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink swiftly so it no longer manages the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("assume-yes", "y", false, "Disable confirmation prompts by assuming 'yes'")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose reporting from swiftly")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help information")
	rootCmd.Flags().Bool("version", false, "Show swiftly version")

	initCmd.Flags().BoolP("no-modify-profile", "n", false, "Do not modify the profile file to set environment variables")
	initCmd.Flags().BoolP("overwrite", "o", false, "Overwrite the existing swiftly installation")
	initCmd.Flags().String("platform", "", "Specify the current Linux platform for swiftly")
	initCmd.Flags().Bool("quiet-shell-followup", false, "Quiet shell follow up commands")
	initCmd.Flags().Bool("skip-install", false, "Skip installing the latest toolchain")

	installCmd.Flags().String("format", "text", "Output format")
	installCmd.Flags().Bool("no-verify", false, "Do not verify the toolchain's PGP signature before installation")
	installCmd.Flags().String("post-install-file", "", "File path to a location for a post installation script")
	installCmd.Flags().String("progress-file", "", "File path where progress information will be written in JSONL format")
	installCmd.Flags().BoolP("use", "u", false, "Mark the newly installed toolchain as in-use")
	installCmd.Flags().Bool("verify", true, "Verify the toolchain's PGP signature before installation")

	listAvailableCmd.Flags().String("format", "text", "Output format")

	useCmd.Flags().String("format", "text", "Output format")
	useCmd.Flags().BoolP("global-default", "g", false, "Set the global default toolchain")
	useCmd.Flags().BoolP("print-location", "p", false, "Print the location of the in-use toolchain")

	listCmd.Flags().String("format", "text", "Output format")

	updateCmd.Flags().Bool("no-verify", false, "Do not verify the toolchain's PGP signature before installation")
	updateCmd.Flags().String("post-install-file", "", "File path to a location for a post installation script")
	updateCmd.Flags().Bool("verify", true, "Verify the toolchain's PGP signature before installation")

	rootCmd.AddCommand(
		installCmd,
		listAvailableCmd,
		useCmd,
		uninstallCmd,
		listCmd,
		updateCmd,
		initCmd,
		selfUpdateCmd,
		runCmd,
		selfUninstallCmd,
		linkCmd,
		unlinkCmd,
	)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"platform": actionLinuxPlatforms(),
	})

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"format":            actionOutputFormats(),
		"post-install-file": carapace.ActionFiles(),
		"progress-file":     carapace.ActionFiles(),
	})
	carapace.Gen(installCmd).PositionalCompletion(actionToolchainSelectors())

	carapace.Gen(listAvailableCmd).FlagCompletion(carapace.ActionMap{
		"format": actionOutputFormats(),
	})
	carapace.Gen(listAvailableCmd).PositionalCompletion(actionToolchainSelectors())

	carapace.Gen(useCmd).FlagCompletion(carapace.ActionMap{
		"format": actionOutputFormats(),
	})
	carapace.Gen(useCmd).PositionalCompletion(actionToolchainSelectors())

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		carapace.Batch(
			actionToolchainSelectors(),
			carapace.ActionValuesDescribed("all", "all installed toolchains"),
		).ToA(),
	)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": actionOutputFormats(),
	})
	carapace.Gen(listCmd).PositionalCompletion(actionToolchainSelectors())

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"post-install-file": carapace.ActionFiles(),
	})
	carapace.Gen(updateCmd).PositionalCompletion(actionToolchainSelectors())

	carapace.Gen(runCmd).PositionalAnyCompletion(carapace.ActionFiles())
}

func actionOutputFormats() carapace.Action {
	return carapace.ActionValues("text", "json").StyleF(style.ForKeyword)
}

func actionLinuxPlatforms() carapace.Action {
	return carapace.ActionValues(
		"amazonlinux2",
		"debian12",
		"fedora39",
		"fedora41",
		"ubi9",
		"ubuntu18.04",
		"ubuntu24.04",
		"ubuntu22.04",
		"ubuntu20.04",
	).StyleF(style.ForKeyword)
}

func actionToolchainSelectors() carapace.Action {
	return carapace.ActionValuesDescribed(
		"latest", "latest stable Swift release",
		"main-snapshot", "latest main development snapshot",
		"6.2", "latest Swift 6.2 patch release",
		"6.1", "latest Swift 6.1 patch release",
		"6.0", "latest Swift 6.0 patch release",
		"5.10", "latest Swift 5.10 patch release",
		"xcode", "currently selected Xcode toolchain on macOS",
	)
}
