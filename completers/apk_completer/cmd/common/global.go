package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddGlobalFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("allow-untrusted", false, "Install packages with untrusted signature or no signature")
	cmd.Flags().String("arch", "", "Temporarily override architecture")
	cmd.Flags().String("cache-dir", "", "Temporarily override the cache directory")
	cmd.Flags().String("cache-max-age", "", "Maximum AGE (in minutes) for index in cache before it's refreshed")
	cmd.Flags().BoolP("force", "f", false, "Enable selected --force-* options (deprecated)")
	cmd.Flags().Bool("force-binary-stdout", false, "Continue even if binary data will be printed to the terminal")
	cmd.Flags().Bool("force-broken-world", false, "DANGEROUS: Delete world constraints until a solution without conflicts is found")
	cmd.Flags().Bool("force-missing-repositories", false, "Continue even if some of the repository indexes are not available")
	cmd.Flags().Bool("force-non-repository", false, "Continue even if packages may be lost on reboot")
	cmd.Flags().Bool("force-old-apk", false, "Continue even if packages use unsupported features")
	cmd.Flags().Bool("force-overwrite", false, "Overwrite files in other packages")
	cmd.Flags().Bool("force-refresh", false, "Do not use cached files (local or from proxy)")
	cmd.Flags().BoolP("interactive", "i", false, "Ask confirmation before performing certain operations")
	cmd.Flags().String("keys-dir", "", "Override directory of trusted keys")
	cmd.Flags().Bool("no-cache", false, "Do not use any local cache path")
	cmd.Flags().Bool("no-check-certificate", false, "Do not validate the HTTPS server certificates")
	cmd.Flags().Bool("no-interactive", false, "Disable interactive mode")
	cmd.Flags().Bool("no-network", false, "Do not use the network")
	cmd.Flags().Bool("no-progress", false, "Disable progress bar even for TTYs")
	cmd.Flags().Bool("print-arch", false, "Print default arch and exit")
	cmd.Flags().Bool("progress", false, "Show progress")
	cmd.Flags().String("progress-fd", "", "Write progress to the specified file descriptor")
	cmd.Flags().Bool("purge", false, "Purge modified configuration and cached packages")
	cmd.Flags().BoolP("quiet", "q", false, "Print less information")
	cmd.Flags().String("repositories-file", "", "Override system repositories")
	cmd.Flags().StringP("repository", "X", "", "Specify additional package repository")
	cmd.Flags().StringP("root", "p", "", "Manage file system at ROOT")
	cmd.Flags().String("timeout", "", "Timeout network connections if no progress is made in TIME seconds")
	cmd.Flags().BoolP("update-cache", "U", false, "Alias for '--cache-max-age 1'")
	cmd.Flags().BoolP("verbose", "v", false, "Print more information (can be specified twice)")
	cmd.Flags().BoolP("version", "V", false, "Print program version and exit")
	cmd.Flags().String("wait", "", "Wait for TIME seconds to get an exclusive repository lock before failing")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"arch":              carapace.ActionValues(), // TODO
		"cache-dir":         carapace.ActionDirectories(),
		"keys-dir":          carapace.ActionDirectories(),
		"progress-fd":       carapace.ActionValues(), // TODO
		"repositories-file": carapace.ActionFiles(),
		"repository":        carapace.ActionValues(), // TODO
		"root":              carapace.ActionDirectories(),
	})
}
