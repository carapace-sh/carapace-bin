package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Executes a dependency graph query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	queryCmd.Flags().Bool("default_visibility", false, "Default visibility for packages that don't set it explicitly ('public' or 'private').")
	queryCmd.Flags().Bool("deleted_packages", false, "A comma-separated list of names of packages which the build system will consider non-existent, even if they are visible somewhere on the package path.")
	queryCmd.Flags().Bool("experimental_check_output_files", false, "Check for modifications made to the output files of a build. Consider setting this flag to false if you don't expect these files to change outside of bazel since it will speed up subsequent runs as they won't have to check a previous run's cache.")
	queryCmd.Flags().Bool("experimental_graphless_query", false, "If true, uses a Query implementation that does not make a copy of the graph. The new implementation only supports --order_output=no, as well as only a subset of output formatters.")
	queryCmd.Flags().Bool("experimental_max_directories_to_eagerly_visit_in_globbing", false, "If non-negative, the first time a glob is evaluated in a package, the subdirectories of the package will be traversed in order to warm filesystem caches and compensate for lack of parallelism in globbing. At most this many directories will be visited.")
	queryCmd.Flags().Bool("fetch", false, "Allows the command to fetch external dependencies. If set to false, the command will utilize any cached version of the dependency, and if none exists, the command will result in failure.")
	queryCmd.Flags().Bool("graph:conditional_edges_limit", false, "The maximum number of condition labels to show. -1 means no truncation and 0 means no annotation. This option is only applicable to --output=graph.")
	queryCmd.Flags().Bool("incompatible_config_setting_private_default_visibility", false, "If incompatible_enforce_config_setting_visibility=false, this is a noop. Else, if this flag is false, any config_setting without an explicit visibility attribute is //visibility:public. If this flag is true, config_setting follows the same visibility logic as all other rules. See https://github.com/bazelbuild/bazel/issues/12933.")
	queryCmd.Flags().Bool("incompatible_enforce_config_setting_visibility", false, "If true, enforce config_setting visibility restrictions. If false, every config_setting is visible to every target. See https://github.com/bazelbuild/bazel/issues/12932.")
	queryCmd.Flags().Bool("incompatible_lexicographical_output", false, "If this option is set, sorts --order_output=auto output in lexicographical order.")
	queryCmd.Flags().BoolP("keep_going", "k", false, "Continue as much as possible after an error.  While the target that failed and those that depend on it cannot be analyzed, other prerequisites of these targets can be.")
	queryCmd.Flags().Bool("legacy_globbing_threads", false, "Number of threads to use for glob evaluation. Takes an integer, or a keyword (\"auto\", \"HOST_CPUS\", \"HOST_RAM\"), optionally followed by an operation ([-|*]<float>) eg. \"auto\", \"HOST_CPUS*.5\". \"auto\" means to use a reasonable value derived from the machine's hardware profile (e.g. the number of processors).")
	queryCmd.Flags().Bool("loading_phase_threads", false, "Number of parallel threads to use for the loading/analysis phase.Takes an integer, or a keyword (\"auto\", \"HOST_CPUS\", \"HOST_RAM\"), optionally followed by an operation ([-|*]<float>) eg. \"auto\", \"HOST_CPUS*.5\". \"auto\" sets a reasonable default based on host resources. Must be at least 1.")
	queryCmd.Flags().Bool("noexperimental_check_output_files", false, "Check for modifications made to the output files of a build. Consider setting this flag to false if you don't expect these files to change outside of bazel since it will speed up subsequent runs as they won't have to check a previous run's cache.")
	queryCmd.Flags().Bool("nofetch", false, "Allows the command to fetch external dependencies. If set to false, the command will utilize any cached version of the dependency, and if none exists, the command will result in failure.")
	queryCmd.Flags().Bool("noincompatible_config_setting_private_default_visibility", false, "If incompatible_enforce_config_setting_visibility=false, this is a noop. Else, if this flag is false, any config_setting without an explicit visibility attribute is //visibility:public. If this flag is true, config_setting follows the same visibility logic as all other rules. See https://github.com/bazelbuild/bazel/issues/12933.")
	queryCmd.Flags().Bool("noincompatible_enforce_config_setting_visibility", false, "If true, enforce config_setting visibility restrictions. If false, every config_setting is visible to every target. See https://github.com/bazelbuild/bazel/issues/12932.")
	queryCmd.Flags().Bool("noincompatible_lexicographical_output", false, "If this option is set, sorts --order_output=auto output in lexicographical order.")
	queryCmd.Flags().Bool("nokeep_going", false, "Continue as much as possible after an error.  While the target that failed and those that depend on it cannot be analyzed, other prerequisites of these targets can be.")
	queryCmd.Flags().Bool("noorder_results", false, "Output the results in dependency-ordered (default) or unordered fashion. The unordered output is faster but only supported when --output is not minrank, maxrank, or graph.")
	queryCmd.Flags().Bool("noshow_loading_progress", false, "If enabled, causes Bazel to print \"Loading package:\" messages.")
	queryCmd.Flags().Bool("nostrict_test_suite", false, "If true, the tests() expression gives an error if it encounters a test_suite containing non-test targets.")
	queryCmd.Flags().Bool("noxml:default_values", false, "If true, rule attributes whose value is not explicitly specified in the BUILD file are printed; otherwise they are omitted.")
	queryCmd.Flags().Bool("noxml:line_numbers", false, "If true, XML output contains line numbers. Disabling this option may make diffs easier to read.  This option is only applicable to --output=xml.")
	queryCmd.Flags().Bool("null", false, "Whether each format is terminated with \\0 instead of newline.")
	queryCmd.Flags().Bool("order_output", false, "Output the results unordered (no), dependency-ordered (deps), or fully ordered (full). The default is 'auto', meaning that results are output either dependency-ordered or fully ordered, depending on the output formatter (dependency-ordered for proto, minrank, maxrank, and graph, fully ordered for all others). When output is fully ordered, nodes are printed in a fully deterministic (total) order. First, all nodes are sorted alphabetically. Then, each node in the list is used as the start of a post-order depth-first search in which outgoing edges to unvisited nodes are traversed in alphabetical order of the successor nodes. Finally, nodes are printed in the reverse of the order in which they were visited.")
	queryCmd.Flags().Bool("order_results", false, "Output the results in dependency-ordered (default) or unordered fashion. The unordered output is faster but only supported when --output is not minrank, maxrank, or graph.")
	queryCmd.Flags().Bool("output", false, "The format in which the query results should be printed. Allowed values for query are: build, graph, streamed_jsonproto, label, label_kind, location, maxrank, minrank, package, proto, streamed_proto, xml.")
	queryCmd.Flags().Bool("package_path", false, "A colon-separated list of where to look for packages. Elements beginning with '%workspace%' are relative to the enclosing workspace. If omitted or empty, the default is the output of 'bazel info default-package-path'.")
	queryCmd.Flags().Bool("show_loading_progress", false, "If enabled, causes Bazel to print \"Loading package:\" messages.")
	queryCmd.Flags().Bool("strict_test_suite", false, "If true, the tests() expression gives an error if it encounters a test_suite containing non-test targets.")
	queryCmd.Flags().Bool("xml:default_values", false, "If true, rule attributes whose value is not explicitly specified in the BUILD file are printed; otherwise they are omitted.")
	queryCmd.Flags().Bool("xml:line_numbers", false, "If true, XML output contains line numbers. Disabling this option may make diffs easier to read.  This option is only applicable to --output=xml.")
	rootCmd.AddCommand(queryCmd)
}
