package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Build a conda package and publish it to a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	publishCmd.Flags().StringP("build-dir", "b", "", "The directory to use for incremental builds artifacts")
	publishCmd.Flags().String("build-number", "", "An optional override for the package's build number")
	publishCmd.Flags().String("build-platform", "linux-64", "The build platform to use for building (defaults to the current platform)")
	publishCmd.Flags().String("build-string-prefix", "", "An optional prefix prepended to the auto-generated build string")
	publishCmd.Flags().BoolP("clean", "c", false, "Whether to clean the build directory before building")
	publishCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	publishCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	publishCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	publishCmd.Flags().Bool("force", false, "Force overwrite existing packages")
	publishCmd.Flags().Bool("generate-attestation", false, "Generate sigstore attestation (prefix.dev only)")
	publishCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	publishCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	publishCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	publishCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	publishCmd.Flags().String("package-format", "", "Archive format and optional compression level, e.g. `conda`, `tar-bz2`, `conda:max`, `conda:15`, `tar-bz2:9`. Numeric ranges match rattler-build: -7..=22 for `.conda`, 1..=9 for `.tar.bz2`")
	publishCmd.Flags().String("path", "", "The path to a directory containing a package manifest, or to a specific manifest file")
	publishCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	publishCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	publishCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	publishCmd.Flags().String("skip-existing", "true", "Skip uploading packages that already exist at the target. This is enabled by default. Use `--no-skip-existing` to disable")
	publishCmd.Flags().String("target-channel", "", "The target channel to publish packages to. Accepts a URL (prefix.dev, anaconda.org, cloudsmith://, s3://, quetz://, artifactory://) or a local filesystem path / `file://` URL for an indexed local channel")
	publishCmd.Flags().String("target-dir", "", "The local filesystem path to copy the built package(s) into (no channel indexing)")
	publishCmd.Flags().StringP("target-platform", "t", "linux-64", "The target platform to build for (defaults to the current platform)")
	publishCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	publishCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	publishCmd.Flags().String("to", "", "The target channel to publish packages to. Accepts a URL (prefix.dev, anaconda.org, cloudsmith://, s3://, quetz://, artifactory://) or a local filesystem path / `file://` URL for an indexed local channel")
	publishCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	publishCmd.Flags().StringSlice("variant", nil, "Override a build variant key with one or more values")
	publishCmd.Flags().StringSliceP("variant-config", "m", nil, "Path to an additional variant configuration YAML file")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"build-dir":             carapace.ActionFiles(),
		"config-file":           carapace.ActionFiles(),
		"path":                  carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
		"skip-existing":         carapace.ActionValues("true", "false"),
		"target-dir":            carapace.ActionFiles(),
		"variant-config":        carapace.ActionFiles(),
	})
}
