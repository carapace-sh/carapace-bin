package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/common"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Build and publish the crate as python packages to pypi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().String("auditwheel", "", "Audit wheel for manylinux compliance")
	publishCmd.Flags().StringP("bindings", "b", "", "Which kind of bindings to use")
	publishCmd.Flags().StringArray("compatibility", nil, "Control the platform tag and PyPI compatibility")
	publishCmd.Flags().Bool("compression-enable-large-file-support", false, "Whether to use large file support for ZIP files")
	publishCmd.Flags().String("compression-level", "", "Zip compression level. Defaults to method default")
	publishCmd.Flags().String("compression-method", "deflated", "Zip compression method")
	publishCmd.Flags().Bool("debug", false, "Do not pass --release to cargo")
	publishCmd.Flags().BoolP("find-interpreter", "f", false, "Find interpreters from the host machine")
	publishCmd.Flags().Bool("generate-stubs", false, "Auto generate Python type stubs by introspecting the binary")
	publishCmd.Flags().Bool("include-debuginfo", false, "Include debug info files in the wheel")
	publishCmd.Flags().StringArrayP("interpreter", "i", nil, "The python versions to build wheels for, given as the executables of interpreters")
	publishCmd.Flags().Bool("no-sdist", false, "Don't build a source distribution")
	publishCmd.Flags().Bool("no-strip", false, "Do not strip the library for minimum file size")
	publishCmd.Flags().Bool("non-interactive", false, "Do not interactively prompt for username/password if the required credentials are missing")
	publishCmd.Flags().StringP("out", "o", "", "The directory to store the built wheels in")
	publishCmd.Flags().StringP("password", "p", "", "Password for pypi or your custom registry")
	publishCmd.Flags().Bool("pgo", false, "Build with Profile-Guided Optimization (PGO)")
	publishCmd.Flags().StringP("repository", "r", "pypi", "The repository (package index) to upload the package to")
	publishCmd.Flags().String("repository-url", "", "The URL of the registry where the wheels are uploaded to")
	publishCmd.Flags().StringArray("sbom-include", nil, "Additional SBOM files to include in the `.dist-info/sboms` directory")
	publishCmd.Flags().Bool("skip-existing", false, "Continue uploading files if one already exists")
	publishCmd.Flags().StringP("username", "u", "", "Username for pypi or your custom registry")
	publishCmd.Flags().Bool("zig", false, "For manylinux targets, use zig to ensure compliance for the chosen manylinux version")
	common.AddCargoFlags(publishCmd)
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"auditwheel":         action.ActionAuditwheel(),
		"bindings":           action.ActionBindings(),
		"compatibility":      action.ActionCompatibility(),
		"compression-method": action.ActionCompressionMethods(),
		"interpreter":        carapace.ActionFiles(),
		"out":                carapace.ActionDirectories(),
		"repository":         action.ActionRepositories(),
		"sbom-include":       carapace.ActionFiles(),
	})
}
