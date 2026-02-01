package golang

import "github.com/carapace-sh/carapace"

// ActionExperiments completes toolchain experiments
//
//	cgocheck2 (enables an expensive cgo rule checker)
//	dwarf5 (enables DWARF version 5 debug info generation)
func ActionExperiments() carapace.Action {
	return carapace.ActionValuesDescribed(
		"arenas", "causes the \"arena\" standard library package to be visible",
		"boringcrypto", "",
		"cgocheck2", "enables an expensive cgo rule checker",
		"dwarf5", "enables DWARF version 5 debug info generation",
		"fieldtrack", "",
		"goroutineleakprofile", "enables the collection of goroutine leak profiles",
		"greenteagc", "enables the Green Tea GC implementation",
		"heapminimum512kib", "reduces the minimum heap size to 512 KiB",
		"jsonv2", "enables the json/v2 package",
		"loopvar", "changes loop semantics so that each iteration gets its own copy",
		"newinliner", "enables a new+improved version of the function inlining phase",
		"preemptibleloops", "",
		"randomizedheapbase64", "enables heap base address randomization on 64-bit platforms",
		"regabiargs", "enables register arguments/results in all compiled Go functions",
		"regabi", "enables all working regabi subexperiments",
		"regabiwrappers", "enables ABI wrappers for calling between ABI0 and ABIInternal functions",
		"runtimefreegc", "enables the runtime to free and reuse memory more eagerly",
		"runtimesecret", "enables the runtime/secret package",
		"simd", "enables the simd package and the compiler's handling of SIMD intrinsics",
		"sizespecializedmalloc", "enables malloc implementations that are specialized per size class",
		"staticlockranking", "",
	)
}
