package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-compile",
	Short: "compiles a single Go packag",
	Long:  "https://pkg.go.dev/cmd/compile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("%", "%", false, "debug non-static initializers")
	rootCmd.Flags().BoolS("+", "+", false, "compiling runtime")
	rootCmd.Flags().BoolS("B", "B", false, "disable bounds checking")
	rootCmd.Flags().BoolS("C", "C", false, "disable printing of columns in error messages")
	rootCmd.Flags().StringS("D", "D", "", "set relative path for local imports")
	rootCmd.Flags().BoolS("E", "E", false, "debug symbol export")
	rootCmd.Flags().StringS("I", "I", "", "add directory to import search path")
	rootCmd.Flags().BoolS("K", "K", false, "debug missing line numbers")
	rootCmd.Flags().BoolS("L", "L", false, "also show actual source file names in error messages for positions affected by //line directives")
	rootCmd.Flags().BoolS("N", "N", false, "disable optimizations")
	rootCmd.Flags().BoolS("S", "S", false, "print assembly listing")
	rootCmd.Flags().BoolS("V", "V", false, "print version and exit")
	rootCmd.Flags().BoolS("W", "W", false, "debug parse tree after type checking")
	rootCmd.Flags().BoolS("asan", "asan", false, "build code compatible with C/C++ address sanitizer")
	rootCmd.Flags().StringS("asmhdr", "asmhdr", "", "write assembly header to file")
	rootCmd.Flags().StringS("bench", "bench", "", "append benchmark times to file")
	rootCmd.Flags().StringS("blockprofile", "blockprofile", "", "write block profile to file")
	rootCmd.Flags().StringS("buildid", "buildid", "", "record id as the build id in the export metadata")
	rootCmd.Flags().StringS("c", "c", "", "concurrency during compilation (1 means no concurrency) (default 12)")
	rootCmd.Flags().BoolS("clobberdead", "clobberdead", false, "clobber dead stack slots (for debugging)")
	rootCmd.Flags().BoolS("clobberdeadreg", "clobberdeadreg", false, "clobber dead registers (for debugging)")
	rootCmd.Flags().BoolS("complete", "complete", false, "compiling complete package (no C or assembly)")
	rootCmd.Flags().StringS("coveragecfg", "coveragecfg", "", "read coverage configuration from file")
	rootCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "write cpu profile to file")
	rootCmd.Flags().StringS("d", "d", "", "enable debugging settings; try -d help")
	rootCmd.Flags().BoolS("dwarf", "dwarf", false, "generate DWARF symbols (default true)")
	rootCmd.Flags().BoolS("dwarfbasentries", "dwarfbasentries", false, "use base address selection entries in DWARF (default true)")
	rootCmd.Flags().BoolS("dwarflocationlists", "dwarflocationlists", false, "add location lists to DWARF in optimized mode (default true)")
	rootCmd.Flags().BoolS("dynlink", "dynlink", false, "support references to Go symbols defined in other shared libraries")
	rootCmd.Flags().BoolS("e", "e", false, "no limit on number of errors reported")
	rootCmd.Flags().StringS("embedcfg", "embedcfg", "", "read go:embed configuration from file")
	rootCmd.Flags().StringS("gendwarfinl", "gendwarfinl", "", "generate DWARF inline info records (default 2)")
	rootCmd.Flags().StringS("goversion", "goversion", "", "required version of the runtime")
	rootCmd.Flags().BoolS("h", "h", false, "halt on error")
	rootCmd.Flags().StringS("importcfg", "importcfg", "", "read import configuration from file")
	rootCmd.Flags().StringS("installsuffix", "installsuffix", "", "set pkg directory suffix")
	rootCmd.Flags().BoolS("j", "j", false, "debug runtime-initialized variables")
	rootCmd.Flags().StringS("json", "json", "", "version,file for JSON compiler/optimizer detail output")
	rootCmd.Flags().BoolS("l", "l", false, "disable inlining -lang string Go language version source code expects")
	rootCmd.Flags().StringS("linkobj", "linkobj", "", "write linker-specific object to file")
	rootCmd.Flags().BoolS("linkshared", "linkshared", false, "generate code that will be linked against Go shared libraries")
	rootCmd.Flags().BoolS("live", "live", false, "debug liveness analysis")
	rootCmd.Flags().BoolS("m", "m", false, "print optimization decisions")
	rootCmd.Flags().StringS("memprofile", "memprofile", "", "write memory profile to file")
	rootCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "set runtime.MemProfileRate to rate")
	rootCmd.Flags().BoolS("msan", "msan", false, "build code compatible with C/C++ memory sanitizer")
	rootCmd.Flags().StringS("mutexprofile", "mutexprofile", "", "write mutex profile to file")
	rootCmd.Flags().BoolS("nolocalimports", "nolocalimports", false, "reject local (relative) imports")
	rootCmd.Flags().StringS("o", "o", "", "write output to file")
	rootCmd.Flags().BoolS("oldcomparable", "oldcomparable", false, "enable old comparable semantics")
	rootCmd.Flags().StringS("p", "p", "", "set expected package import path")
	rootCmd.Flags().BoolS("pack", "pack", false, "write to file.a instead of file.o")
	rootCmd.Flags().StringS("pgoprofile", "pgoprofile", "", "read profile from file")
	rootCmd.Flags().BoolS("r", "r", false, "debug generated wrappers")
	rootCmd.Flags().BoolS("race", "race", false, "enable race detector")
	rootCmd.Flags().BoolS("shared", "shared", false, "generate code that can be linked into a shared library")
	rootCmd.Flags().BoolS("smallframes", "smallframes", false, "reduce the size limit for stack allocated objects")
	rootCmd.Flags().StringS("spectre", "spectre", "", "enable spectre mitigations in list (all, index, ret)")
	rootCmd.Flags().BoolS("std", "std", false, "compiling standard library")
	rootCmd.Flags().StringS("symabis", "symabis", "", "read symbol ABIs from file")
	rootCmd.Flags().BoolS("t", "t", false, "enable tracing for debugging the compiler")
	rootCmd.Flags().StringS("traceprofile", "traceprofile", "", "write an execution trace to file")
	rootCmd.Flags().StringS("trimpath", "trimpath", "", "remove prefix from recorded source file paths")
	rootCmd.Flags().BoolS("v", "v", false, "increase debug verbosity")
	rootCmd.Flags().BoolS("w", "w", false, "debug type checking")
	rootCmd.Flags().BoolS("wb", "wb", false, "enable write barrier (default true)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"I":            carapace.ActionDirectories(),
		"asmhdr":       carapace.ActionFiles(),
		"bench":        carapace.ActionFiles(),
		"blockprofile": carapace.ActionFiles(),
		"coveragecfg":  carapace.ActionFiles(),
		"cpuprofile":   carapace.ActionFiles(),
		"embedcfg":     carapace.ActionFiles(),
		"goversion":    golang.ActionVersions(),
		"linkobj":      carapace.ActionFiles(),
		"memprofile":   carapace.ActionFiles(),
		"mutexprofile": carapace.ActionFiles(),
		"o":            carapace.ActionFiles(),
		"pgoprofile":   carapace.ActionFiles(),
		"spectre":      carapace.ActionValues("all", "index", "ret").UniqueList(","),
		"traceprofile": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}
