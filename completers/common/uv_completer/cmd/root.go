package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uv",
	Short: "An extremely fast Python package manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringSlice("allow-insecure-host", nil, "Allow insecure connections to a host")
	rootCmd.PersistentFlags().Bool("allow-python-downloads", false, "Allow automatically downloading Python when required. [env: \"UV_PYTHON_DOWNLOADS=auto\"]")
	rootCmd.PersistentFlags().String("cache-dir", "", "Path to the cache directory")
	rootCmd.PersistentFlags().String("color", "", "Control the use of color in output")
	rootCmd.PersistentFlags().String("config-file", "", "The path to a `uv.toml` file to use for configuration")
	rootCmd.PersistentFlags().String("directory", "", "Change to the given directory prior to running the command")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display the concise help for this command")
	rootCmd.PersistentFlags().Bool("isolated", false, "Avoid discovering a `pyproject.toml` or `uv.toml` file [env: UV_ISOLATED=]")
	rootCmd.PersistentFlags().Bool("managed-python", false, "Require use of uv-managed Python versions [env: UV_MANAGED_PYTHON=]")
	rootCmd.PersistentFlags().Bool("native-tls", false, "(Deprecated: use `--system-certs` instead.) Whether to load TLS certificates from the platform's native certificate store [env: UV_NATIVE_TLS=]")
	rootCmd.PersistentFlags().BoolP("no-cache", "n", false, "Avoid reading from or writing to the cache, instead using a temporary directory for the duration of the operation")
	rootCmd.PersistentFlags().Bool("no-cache-dir", false, "Avoid reading from or writing to the cache, instead using a temporary directory for the duration of the operation")
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable colors")
	rootCmd.PersistentFlags().Bool("no-config", false, "Avoid discovering configuration files (`pyproject.toml`, `uv.toml`)")
	rootCmd.PersistentFlags().Bool("no-installer-metadata", false, "Skip writing `uv` installer metadata files (e.g., `INSTALLER`, `REQUESTED`, and `direct_url.json`) to site-packages `.dist-info` directories [env: UV_NO_INSTALLER_METADATA=]")
	rootCmd.PersistentFlags().Bool("no-managed-python", false, "Disable use of uv-managed Python versions [env: UV_NO_MANAGED_PYTHON=]")
	rootCmd.PersistentFlags().Bool("no-native-tls", false, "")
	rootCmd.PersistentFlags().Bool("no-offline", false, "")
	rootCmd.PersistentFlags().Bool("no-preview", false, "")
	rootCmd.PersistentFlags().Bool("no-progress", false, "Hide all progress outputs [env: UV_NO_PROGRESS=]")
	rootCmd.PersistentFlags().Bool("no-python-downloads", false, "Disable automatic downloads of Python. [env: \"UV_PYTHON_DOWNLOADS=never\"]")
	rootCmd.PersistentFlags().Bool("no-system-certs", false, "")
	rootCmd.PersistentFlags().Bool("offline", false, "Disable network access [env: UV_OFFLINE=]")
	rootCmd.PersistentFlags().Bool("preview", false, "Whether to enable all experimental preview features [env: UV_PREVIEW=]")
	rootCmd.PersistentFlags().StringSlice("preview-feature", nil, "Enable experimental preview features")
	rootCmd.PersistentFlags().StringSlice("preview-features", nil, "Enable experimental preview features")
	rootCmd.PersistentFlags().String("project", "", "Discover a project in the given directory")
	rootCmd.PersistentFlags().String("python-fetch", "", "Deprecated version of [`Self::python_downloads`]")
	rootCmd.PersistentFlags().String("python-preference", "", "")
	rootCmd.PersistentFlags().CountP("quiet", "q", "Use quiet output")
	rootCmd.PersistentFlags().Bool("show-settings", false, "Show the resolved settings for the current command")
	rootCmd.PersistentFlags().Bool("system-certs", false, "Whether to load TLS certificates from the platform's native certificate store [env: UV_SYSTEM_CERTS=]")
	rootCmd.PersistentFlags().StringSlice("trusted-host", nil, "Allow insecure connections to a host")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Display the uv version")
	rootCmd.Flag("allow-python-downloads").Hidden = true
	rootCmd.Flag("isolated").Hidden = true
	rootCmd.Flag("native-tls").Hidden = true
	rootCmd.Flag("no-cache-dir").Hidden = true
	rootCmd.Flag("no-color").Hidden = true
	rootCmd.Flag("no-installer-metadata").Hidden = true
	rootCmd.Flag("no-native-tls").Hidden = true
	rootCmd.Flag("no-offline").Hidden = true
	rootCmd.Flag("no-preview").Hidden = true
	rootCmd.Flag("no-system-certs").Hidden = true
	rootCmd.Flag("preview").Hidden = true
	rootCmd.Flag("preview-feature").Hidden = true
	rootCmd.Flag("preview-features").Hidden = true
	rootCmd.Flag("python-fetch").Hidden = true
	rootCmd.Flag("python-preference").Hidden = true
	rootCmd.Flag("show-settings").Hidden = true
	rootCmd.Flag("trusted-host").Hidden = true
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"color": carapace.ActionValuesDescribed(
			"auto", "Enables colored output only when the output is going to a terminal or TTY with support",
			"always", "Enables colored output regardless of the detected environment",
			"never", "Disables colored output",
		),
		"config-file": carapace.ActionFiles(),
		"directory":   carapace.ActionDirectories(),
		"project":     carapace.ActionDirectories(),
		"python-fetch": carapace.ActionValuesDescribed(
			"automatic", "Automatically download managed Python installations when needed",
			"manual", "Do not automatically download managed Python installations; require explicit installation",
			"never", "Do not ever allow Python downloads",
		),
		"python-preference": carapace.ActionValuesDescribed(
			"only-managed", "Only use managed Python installations; never use system Python installations",
			"managed", "Prefer managed Python installations over system Python installations",
			"system", "Prefer system Python installations over managed Python installations",
			"only-system", "Only use system Python installations; never use managed Python installations",
		),
	})
}
