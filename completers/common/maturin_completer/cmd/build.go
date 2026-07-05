package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/common"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the crate into python packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("auditwheel", "", "Audit wheel for manylinux compliance")
	buildCmd.Flags().StringP("bindings", "b", "", "Which kind of bindings to use")
	buildCmd.Flags().StringArray("compatibility", nil, "Control the platform tag and PyPI compatibility")
	buildCmd.Flags().Bool("compression-enable-large-file-support", false, "Whether to use large file support for ZIP files")
	buildCmd.Flags().String("compression-level", "", "Zip compression level. Defaults to method default")
	buildCmd.Flags().String("compression-method", "deflated", "Zip compression method")
	buildCmd.Flags().BoolP("find-interpreter", "f", false, "Find interpreters from the host machine")
	buildCmd.Flags().Bool("generate-stubs", false, "Auto generate Python type stubs by introspecting the binary")
	buildCmd.Flags().Bool("include-debuginfo", false, "Include debug info files in the wheel")
	buildCmd.Flags().StringArrayP("interpreter", "i", nil, "The python versions to build wheels for, given as the executables of interpreters")
	buildCmd.Flags().StringP("out", "o", "", "The directory to store the built wheels in")
	buildCmd.Flags().Bool("pgo", false, "Build with Profile-Guided Optimization (PGO)")
	buildCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	buildCmd.Flags().StringArray("sbom-include", nil, "Additional SBOM files to include in the `.dist-info/sboms` directory")
	buildCmd.Flags().Bool("sdist", false, "Build a source distribution and build wheels from it")
	buildCmd.Flags().String("strip", "", "Strip the library for minimum file size")
	buildCmd.Flags().Bool("zig", false, "For manylinux targets, use zig to ensure compliance for the chosen manylinux version")
	common.AddCargoFlags(buildCmd)
	buildCmd.Flag("strip").NoOptDefVal = "true"
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"auditwheel":         action.ActionAuditwheel(),
		"bindings":           action.ActionBindings(),
		"compatibility":      action.ActionCompatibility(),
		"compression-method": action.ActionCompressionMethods(),
		"interpreter":        carapace.ActionFiles(),
		"out":                carapace.ActionDirectories(),
		"sbom-include":       carapace.ActionFiles(),
		"strip":              carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
	})
}
