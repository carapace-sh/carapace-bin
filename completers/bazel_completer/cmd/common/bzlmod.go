package common

import "github.com/spf13/cobra"

func AddBzlmodFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice("allow_yanked_versions", []string{}, "Specified the module versions in the form of `<module1>@<version1>, <module2>@<version2>` that will be allowed in the resolved dependency graph even if they are declared yanked in the registry where they come from (if they are not coming from a NonRegistryOverride)")
	cmd.Flags().String("check_bazel_compatibility", "", "Check bazel version compatibility of Bazel modules")
	cmd.Flags().String("check_direct_dependencies", "", "Check if the direct `bazel_dep` dependencies declared in the root module are the same versions you get in the resolved dependency graph")
	cmd.Flags().Bool("ignore_dev_dependency", false, "If true, Bazel ignores `bazel_dep` and `use_extension` declared as `dev_dependency` in the MODULE.bazel of the root module")
	cmd.Flags().String("lockfile_mode", "", "Specifies how and whether or not to use the lockfile")
	cmd.Flags().Bool("no-ignore_dev_dependency", false, "If true, Bazel ignores `bazel_dep` and `use_extension` declared as `dev_dependency` in the MODULE.bazel of the root module")
	cmd.Flags().StringSlice("override_module", []string{}, "Override a module with a local path in the form of <module name>=<path>")
	cmd.Flags().StringSlice("registry", []string{}, "Specifies the registries to use to locate Bazel module dependencies")
	cmd.Flags().String("vendor_dir", "", "Specifies the directory that should hold the external repositories in vendor mode, whether for the purpose of fetching them into it or using them while building")

	// TODO add flag completion
}
