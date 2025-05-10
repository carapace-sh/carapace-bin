package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-link",
	Short: "go linker",
	Long:  "https://pkg.go.dev/cmd/link",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("B", "B", "", "add an ELF NT_GNU_BUILD_ID note when using ELF")
	rootCmd.Flags().StringS("E", "E", "", "set entry symbol name")
	rootCmd.Flags().StringS("H", "H", "", "set header type")
	rootCmd.Flags().StringS("I", "I", "", "use linker as ELF dynamic linker")
	rootCmd.Flags().StringSliceS("L", "L", nil, "add specified directory to library path")
	rootCmd.Flags().StringS("R", "R", "", "set address rounding quantum (default -1)")
	rootCmd.Flags().StringS("T", "T", "", "set text segment address (default -1)")
	rootCmd.Flags().StringS("X", "X", "", "add string value definition of the form importpath.name=value")
	rootCmd.Flags().BoolS("asan", "asan", false, "enable ASan interface")
	rootCmd.Flags().BoolS("aslr", "aslr", false, "enable ASLR for buildmode=c-shared on windows (default true)")
	rootCmd.Flags().StringS("benchmark", "benchmark", "", "set to 'mem' or 'cpu' to enable phase benchmarking")
	rootCmd.Flags().StringS("benchmarkprofile", "benchmarkprofile", "", "emit phase profiles to base_phase.{cpu,mem}prof")
	rootCmd.Flags().StringS("buildid", "buildid", "", "record id as Go toolchain build id")
	rootCmd.Flags().StringS("buildmode", "buildmode", "", "set build mode")
	rootCmd.Flags().BoolS("c", "c", false, "dump call graph")
	rootCmd.Flags().StringS("capturehostobjs", "capturehostobjs", "", "capture host object files loaded during internal linking to specified dir")
	rootCmd.Flags().BoolS("compressdwarf", "compressdwarf", false, "compress DWARF if possible (default true)")
	rootCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "write cpu profile to file")
	rootCmd.Flags().BoolS("d", "d", false, "disable dynamic executable")
	rootCmd.Flags().BoolS("debugnosplit", "debugnosplit", false, "dump nosplit call graph")
	rootCmd.Flags().StringS("debugtextsize", "debugtextsize", "", "debug text section max size")
	rootCmd.Flags().StringS("debugtramp", "debugtramp", "", "debug trampolines")
	rootCmd.Flags().BoolS("dumpdep", "dumpdep", false, "dump symbol dependency graph")
	rootCmd.Flags().StringS("extar", "extar", "", "archive program for buildmode=c-archive")
	rootCmd.Flags().StringS("extld", "extld", "", "use linker when linking in external mode")
	rootCmd.Flags().StringS("extldflags", "extldflags", "", "pass flags to external linker")
	rootCmd.Flags().BoolS("f", "f", false, "ignore version mismatch")
	rootCmd.Flags().BoolS("g", "g", false, "disable go package data checks")
	rootCmd.Flags().BoolS("h", "h", false, "halt on error")
	rootCmd.Flags().StringS("importcfg", "importcfg", "", "read import configuration from file")
	rootCmd.Flags().StringS("installsuffix", "installsuffix", "", "set package directory suffix")
	rootCmd.Flags().StringS("k", "k", "", "set field tracking symbol")
	rootCmd.Flags().StringS("libgcc", "libgcc", "", "compiler support lib for internal linking")
	rootCmd.Flags().StringS("linkmode", "linkmode", "", "set link mode")
	rootCmd.Flags().BoolS("linkshared", "linkshared", false, "link against installed Go shared libraries")
	rootCmd.Flags().StringS("memprofile", "memprofile", "", "write memory profile to file")
	rootCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "set runtime.MemProfileRate to rate")
	rootCmd.Flags().BoolS("msan", "msan", false, "enable MSan interface")
	rootCmd.Flags().StringS("n", "n", "", "symbol table")
	rootCmd.Flags().StringS("o", "o", "", "write output to file")
	rootCmd.Flags().StringS("pluginpath", "pluginpath", "", "full path name for plugin")
	rootCmd.Flags().StringS("r", "r", "", "set the ELF dynamic linker search path to dir1:dir2:...")
	rootCmd.Flags().BoolS("race", "race", false, "enable race detector")
	rootCmd.Flags().BoolS("s", "s", false, "disable symbol table")
	rootCmd.Flags().StringS("strictdups", "strictdups", "", "sanity check duplicate symbol contents during object file reading (1=warn 2=err)")
	rootCmd.Flags().StringS("tmpdir", "tmpdir", "", "use directory for temporary files")
	rootCmd.Flags().BoolS("v", "v", false, "print link trace")
	rootCmd.Flags().BoolS("w", "w", false, "disable DWARF generation")

	// TODO extldflags completion
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"H": carapace.ActionValues("windowsgui"), // TODO other header types
		"I": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"L":          carapace.ActionDirectories(),
		"benchmark":  carapace.ActionValues("mem", "cpu"),
		"buildmode":  golang.ActionBuildmodes(),
		"cpuprofile": carapace.ActionFiles(),
		"extar": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"extld": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"importcfg": carapace.ActionFiles(),
		"libgcc": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"linkmode":   carapace.ActionValues("internal", "external", "auto").StyleF(style.ForKeyword),
		"memprofile": carapace.ActionFiles(),
		"o":          carapace.ActionFiles(),
		"r":          carapace.ActionDirectories().List(","),
		"tmpdir":     carapace.ActionDirectories(),
	})
}
