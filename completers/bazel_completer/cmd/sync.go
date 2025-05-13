package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs all repositories specified in the workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().Bool("configure", false, "Only sync repositories marked as 'configure' for system-configuration purpose.")
	syncCmd.Flags().Bool("default_visibility", false, "Default visibility for packages that don't set it explicitly ('public' or 'private').")
	syncCmd.Flags().Bool("deleted_packages", false, "A comma-separated list of names of packages which the build system will consider non-existent, even if they are visible somewhere on the package path.")
	syncCmd.Flags().Bool("experimental_check_output_files", false, "Check for modifications made to the output files of a build. Consider setting this flag to false if you don't expect these files to change outside of bazel since it will speed up subsequent runs as they won't have to check a previous run's cache.")
	syncCmd.Flags().Bool("experimental_max_directories_to_eagerly_visit_in_globbing", false, "If non-negative, the first time a glob is evaluated in a package, the subdirectories of the package will be traversed in order to warm filesystem caches and compensate for lack of parallelism in globbing. At most this many directories will be visited.")
	syncCmd.Flags().Bool("fetch", false, "Allows the command to fetch external dependencies. If set to false, the command will utilize any cached version of the dependency, and if none exists, the command will result in failure.")
	syncCmd.Flags().Bool("incompatible_config_setting_private_default_visibility", false, "If incompatible_enforce_config_setting_visibility=false, this is a noop. Else, if this flag is false, any config_setting without an explicit visibility attribute is //visibility:public. If this flag is true, config_setting follows the same visibility logic as all other rules. See https://github.com/bazelbuild/bazel/issues/12933.")
	syncCmd.Flags().Bool("incompatible_enforce_config_setting_visibility", false, "If true, enforce config_setting visibility restrictions. If false, every config_setting is visible to every target. See https://github.com/bazelbuild/bazel/issues/12932.")
	syncCmd.Flags().BoolP("keep_going", "k", false, "Continue as much as possible after an error.  While the target that failed and those that depend on it cannot be analyzed, other prerequisites of these targets can be.")
	syncCmd.Flags().Bool("legacy_globbing_threads", false, "Number of threads to use for glob evaluation. Takes an integer, or a keyword (\"auto\", \"HOST_CPUS\", \"HOST_RAM\"), optionally followed by an operation ([-|*]<float>) eg. \"auto\", \"HOST_CPUS*.5\". \"auto\" means to use a reasonable value derived from the machine's hardware profile (e.g. the number of processors).")
	syncCmd.Flags().Bool("loading_phase_threads", false, "Number of parallel threads to use for the loading/analysis phase.Takes an integer, or a keyword (\"auto\", \"HOST_CPUS\", \"HOST_RAM\"), optionally followed by an operation ([-|*]<float>) eg. \"auto\", \"HOST_CPUS*.5\". \"auto\" sets a reasonable default based on host resources. Must be at least 1.")
	syncCmd.Flags().Bool("noconfigure", false, "Only sync repositories marked as 'configure' for system-configuration purpose.")
	syncCmd.Flags().Bool("noexperimental_check_output_files", false, "Check for modifications made to the output files of a build. Consider setting this flag to false if you don't expect these files to change outside of bazel since it will speed up subsequent runs as they won't have to check a previous run's cache.")
	syncCmd.Flags().Bool("nofetch", false, "Allows the command to fetch external dependencies. If set to false, the command will utilize any cached version of the dependency, and if none exists, the command will result in failure.")
	syncCmd.Flags().Bool("noincompatible_config_setting_private_default_visibility", false, "If incompatible_enforce_config_setting_visibility=false, this is a noop. Else, if this flag is false, any config_setting without an explicit visibility attribute is //visibility:public. If this flag is true, config_setting follows the same visibility logic as all other rules. See https://github.com/bazelbuild/bazel/issues/12933.")
	syncCmd.Flags().Bool("noincompatible_enforce_config_setting_visibility", false, "If true, enforce config_setting visibility restrictions. If false, every config_setting is visible to every target. See https://github.com/bazelbuild/bazel/issues/12932.")
	syncCmd.Flags().Bool("nokeep_going", false, "Continue as much as possible after an error.  While the target that failed and those that depend on it cannot be analyzed, other prerequisites of these targets can be.")
	syncCmd.Flags().Bool("noshow_loading_progress", false, "If enabled, causes Bazel to print \"Loading package:\" messages.")
	syncCmd.Flags().Bool("only", false, "If this option is given, only sync the repositories specified with this option. Still consider all (or all configure-like, of --configure is given) outdated.")
	syncCmd.Flags().Bool("package_path", false, "A colon-separated list of where to look for packages. Elements beginning with '%workspace%' are relative to the enclosing workspace. If omitted or empty, the default is the output of 'bazel info default-package-path'.")
	syncCmd.Flags().Bool("show_loading_progress", false, "If enabled, causes Bazel to print \"Loading package:\" messages.")
	rootCmd.AddCommand(syncCmd)
}
