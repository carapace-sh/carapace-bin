package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func addCommonRebuildFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("dry", "n", false, "Only print actions, without performing them")
	cmd.Flags().BoolP("ask", "a", false, "Ask for confirmation")
	cmd.Flags().Bool("no-nom", false, "Don't use nix-output-monitor for the build process")
	cmd.Flags().StringP("out-link", "o", "", "Path to save the result link")
	cmd.Flags().StringP("diff", "d", "auto", "Whether to display a package diff")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"diff":     carapace.ActionValues("auto", "always", "never"),
		"out-link": carapace.ActionFiles(),
	})

	addNixBuildPassthroughFlags(cmd)
}

func addNixBuildPassthroughFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("max-jobs", "j", "", "Number of concurrent jobs Nix should run")
	cmd.Flags().String("cores", "", "Number of cores Nix should utilize")
	cmd.Flags().String("log-format", "", "Logging format used by Nix")
	cmd.Flags().BoolP("keep-going", "k", false, "Continue building despite encountering errors")
	cmd.Flags().BoolP("keep-failed", "K", false, "Keep build outputs from failed builds")
	cmd.Flags().Bool("fallback", false, "Attempt to build locally if substituters fail")
	cmd.Flags().Bool("repair", false, "Repair corrupted store paths")
	cmd.Flags().String("builders", "", "Explicitly define remote builders")
	cmd.Flags().StringArrayP("include", "I", nil, "Paths to include")
	cmd.Flags().BoolP("print-build-logs", "L", false, "Print build logs directly to stdout")
	cmd.Flags().BoolP("show-trace", "t", false, "Display tracebacks on errors")
	cmd.Flags().Bool("accept-flake-config", false, "Accept configuration from flakes")
	cmd.Flags().Bool("refresh", false, "Refresh flakes to the latest revision")
	cmd.Flags().Bool("impure", false, "Allow impure builds")
	cmd.Flags().Bool("offline", false, "Build without internet access")
	cmd.Flags().Bool("no-net", false, "Prohibit network usage")
	cmd.Flags().Bool("recreate-lock-file", false, "Recreate the flake.lock file entirely")
	cmd.Flags().Bool("no-update-lock-file", false, "Do not update the flake.lock file")
	cmd.Flags().Bool("no-write-lock-file", false, "Do not write a lock file")
	cmd.Flags().Bool("no-use-registries", false, "Do not use registries")
	cmd.Flags().Bool("no-registries", false, "Do not use registries (deprecated)")
	cmd.Flags().Bool("commit-lock-file", false, "Commit the lock file after updates")
	cmd.Flags().BoolP("no-build-output", "Q", false, "Suppress build output")
	cmd.Flags().Bool("use-substitutes", false, "Use substitutes when copying")
	cmd.Flags().Bool("json", false, "Output results in JSON format")
}

func addUpdateFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("update", "u", false, "Update all flake inputs")
	cmd.Flags().StringArrayP("update-input", "U", nil, "Update the specified flake input(s)")
	cmd.MarkFlagsMutuallyExclusive("update", "update-input")
}

func addOsRebuildFlags(cmd *cobra.Command) {
	addCommonRebuildFlags(cmd)
	addUpdateFlags(cmd)
	cmd.Flags().StringP("hostname", "H", "", "Select this hostname from nixosConfigurations")
	cmd.Flags().StringP("specialisation", "s", "", "Explicitly select some specialisation")
	cmd.Flags().BoolP("no-specialisation", "S", false, "Ignore specialisations")
	cmd.Flags().Bool("install-bootloader", false, "Install bootloader for switch and boot commands")
	cmd.Flags().BoolP("bypass-root-check", "R", false, "Don't panic if calling nh as root")
	cmd.Flags().String("target-host", "", "Deploy the built configuration to a different host over SSH")
	cmd.Flags().String("build-host", "", "Build the configuration on a different host over SSH")
	cmd.Flags().Bool("no-validate", false, "Skip pre-activation system validation checks")
}
