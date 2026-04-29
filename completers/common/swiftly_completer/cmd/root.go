package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swiftly",
	Short: "A utility for installing and managing Swift toolchains",
	Long:  "https://www.swift.org/swiftly/documentation/swiftly/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Disable confirmation prompts by assuming yes")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose reporting")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("version", false, "Print version")

	rootCmd.AddCommand(
		newInstallCmd(),
		newListAvailableCmd(),
		newUseCmd(),
		newUninstallCmd(),
		newListCmd(),
		newUpdateCmd(),
		newInitCmd(),
		newSelfUpdateCmd(),
		newRunCmd(),
		newSelfUninstallCmd(),
		newLinkCmd(),
		newUnlinkCmd(),
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"install", "Install a Swift toolchain",
			"list-available", "List toolchains available for install",
			"use", "Select an installed toolchain",
			"uninstall", "Uninstall a toolchain",
			"list", "List installed toolchains",
			"update", "Update an installed toolchain",
			"init", "Initialize swiftly",
			"self-update", "Update swiftly itself",
			"run", "Run a command with the selected toolchain",
			"self-uninstall", "Uninstall swiftly itself",
			"link", "Resume managing the active toolchain",
			"unlink", "Stop managing the active toolchain",
		).StyleF(style.ForKeyword),
	)
}

func addToolchainArg(cmd *cobra.Command) {
	carapace.Gen(cmd).PositionalCompletion(actionInstalledToolchains())
}

func addPlatformFlag(cmd *cobra.Command) {
	cmd.Flags().String("platform", "", "Select platform")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{"platform": carapace.ActionValues("linux", "macos", "ubuntu", "rhel", "amazonlinux")})
}

func newInstallCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "install", Short: "Install a Swift toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().Bool("post-install-file", false, "Use post-install file")
	cmd.Flags().Bool("verify", false, "Verify the toolchain signature")
	cmd.Flags().String("use", "", "Use the installed toolchain")
	addPlatformFlag(cmd)
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{"use": carapace.ActionValues("true", "false")})
	carapace.Gen(cmd).PositionalCompletion(actionAvailableToolchains())
	return cmd
}

func newListAvailableCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "list-available", Short: "List toolchains available for install", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().Bool("show-all", false, "Show all available toolchains")
	cmd.Flags().String("format", "", "Output format")
	addPlatformFlag(cmd)
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{"format": carapace.ActionValues("text", "json")})
	return cmd
}

func newUseCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "use", Short: "Select an installed toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().Bool("global", false, "Use globally")
	cmd.Flags().Bool("local", false, "Use locally")
	addToolchainArg(cmd)
	return cmd
}

func newUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "uninstall", Short: "Uninstall a toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	addToolchainArg(cmd)
	return cmd
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "list", Short: "List installed toolchains", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().String("format", "", "Output format")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{"format": carapace.ActionValues("text", "json")})
	return cmd
}

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "update", Short: "Update an installed toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().Bool("verify", false, "Verify the toolchain signature")
	cmd.Flags().String("to", "", "Target version")
	addToolchainArg(cmd)
	return cmd
}

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "init", Short: "Initialize swiftly", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().BoolP("no-modify-profile", "n", false, "Do not modify shell profile")
	cmd.Flags().BoolP("quiet-shell-followup", "q", false, "Quiet shell follow-up commands")
	cmd.Flags().Bool("skip-install", false, "Skip installing the latest toolchain")
	addPlatformFlag(cmd)
	return cmd
}

func newSelfUpdateCmd() *cobra.Command {
	return &cobra.Command{Use: "self-update", Short: "Update swiftly itself", Run: func(cmd *cobra.Command, args []string) {}}
}
func newSelfUninstallCmd() *cobra.Command {
	return &cobra.Command{Use: "self-uninstall", Short: "Uninstall swiftly itself", Run: func(cmd *cobra.Command, args []string) {}}
}
func newLinkCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "link", Short: "Resume managing the active toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	addToolchainArg(cmd)
	return cmd
}
func newUnlinkCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "unlink", Short: "Stop managing the active toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	addToolchainArg(cmd)
	return cmd
}

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "run", Short: "Run a command with the selected toolchain", Run: func(cmd *cobra.Command, args []string) {}}
	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionExecutables())
	return cmd
}

func actionInstalledToolchains() carapace.Action {
	return carapace.ActionExecCommand("swiftly", "list", "--format", "json")(func(output []byte) carapace.Action {
		return carapace.ActionValues()
	})
}

func actionAvailableToolchains() carapace.Action {
	return carapace.ActionExecCommand("swiftly", "list-available", "--format", "json")(func(output []byte) carapace.Action {
		return carapace.ActionValues()
	})
}
