package rust

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionCodegenOptions completes codegen names and values
func ActionCodegenOptions() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionCodegenOptionNames().Suffix("=")
		default:
			return ActionCodegenOptionValues(c.Parts[0])
		}
	})
}

// ActionCodegenOptionNames completes codegen option names
//
//	ar (this option is deprecated and does nothing)
//	code-model (choose the code model to use)
func ActionCodegenOptionNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ar", "this option is deprecated and does nothing",
		"code-model", "choose the code model to use",
		"codegen-units", "divide crate into N units to optimize in parallel",
		"collapse-macro-debuginfo", "set option to collapse debuginfo for macros",
		"control-flow-guard", "use Windows Control Flow Guard",
		"debug-assertions", "explicitly enable the `cfg(debug_assertions)` directive",
		"debuginfo", "debug info emission level",
		"default-linker-libraries", "allow the linker to link its default libraries",
		"dlltool", "import library generation tool",
		"dwarf-version", "version of DWARF debug information to emit",
		"embed-bitcode", "emit bitcode in rlibs",
		"extra-filename", "extra data to put in each output filename",
		"force-frame-pointers", "force use of the frame pointers",
		"force-unwind-tables", "force use of unwind tables",
		"incremental", "enable incremental compilation",
		"inline-threshold", "this option is deprecated and does nothing",
		"instrument-coverage", "instrument the generated code to support LLVM source-based code coverage reports",
		"link-arg", "a single extra argument to append to the linker invocation",
		"link-args", "extra arguments to append to the linker invocation",
		"link-dead-code", "try to generate and link dead code",
		"link-self-contained", "control whether to link Rust provided C objects/libraries or rely on a C toolchain or linker installed in the system",
		"linker", "system linker to link outputs with",
		"linker-features", "a comma-separated list of linker features to enable (+) or disable (-): `lld`",
		"linker-flavor", "linker flavor",
		"linker-plugin-lto", "generate build artifacts that are compatible with linker-based LTO",
		"llvm-args", "a list of arguments to pass to LLVM",
		"lto", "perform LLVM link-time optimizations",
		"metadata", "metadata to mangle symbol names with",
		"no-prepopulate-passes", "give an empty list of passes to the pass manager",
		"no-redzone", "disable the use of the redzone",
		"no-stack-check", "this option is deprecated and does nothing",
		"no-vectorize-loops", "disable loop vectorization optimization passes",
		"no-vectorize-slp", "disable LLVM's SLP vectorization pass",
		"opt-level", "optimization level",
		"overflow-checks", "use overflow checks for integer arithmetic",
		"panic", "panic strategy to compile crate with",
		"passes", "a list of extra LLVM passes to run",
		"prefer-dynamic", "prefer dynamic linking to static linking",
		"profile-generate", "compile the program with profiling instrumentation",
		"profile-use", "use the given `.profdata` file for profile-guided optimization",
		"relocation-model", "control generation of position-independent code",
		"relro-level", "choose which RELRO level to use",
		"remark", "output remarks for these optimization passes",
		"rpath", "set rpath values in libs/exes",
		"save-temps", "save all temporary output files during compilation",
		"soft-float", "deprecated option: use soft float ABI",
		"split-debuginfo", "how to handle split-debuginfo, a platform-specific option",
		"strip", "tell the linker which information to strip",
		"symbol-mangling-version", "which mangling version to use for symbol names",
		"target-cpu", "select target processor",
		"target-feature", "target specific attributes. This feature is unsafe.",
		"unsafe-allow-abi-mismatch", "Allow incompatible target modifiers in dependency crates",
	).Tag("codegen options").Uid("rust", "codegen-option")
}

// ActionCodegenOptionValues completion codegen option value
func ActionCodegenOptionValues(s string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {

		return map[string]carapace.Action{
			// TODO complete more
			"code-model": carapace.ActionValues("tiny", "small", "kernel", "medium", "large"),
			"strip":      carapace.ActionValues("none", "debuginfo", "symbols").StyleF(style.ForKeyword),
		}[s]
	})
}
