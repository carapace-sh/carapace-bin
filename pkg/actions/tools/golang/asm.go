package golang

import "github.com/carapace-sh/carapace"

// ActionAsmPredefinedSymbolValues completes predefined assembler symbols and their values
//
//	GOOS_linux=1
//	GOARCH_amd64=1
func ActionAsmPredefinedSymbolValues() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionAsmPredefinedSymbols().Suffix("=")
		default:
			return ActionAsmPredefinedValues(c.Parts[0])
		}
	})
}

// ActionAsmPredefinedSymbols completes predefined assembler symbols
//
//	GOOS_linux
//	GOARCH_amd64
//	GOAMD64_v1
func ActionAsmPredefinedSymbols() carapace.Action {
	return carapace.Batch(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			osSymbols := make([]string, 0)
			for _, os := range []string{"aix", "android", "darwin", "dragonfly", "freebsd", "hurd", "illumos", "ios", "js", "linux", "nacl", "netbsd", "openbsd", "plan9", "solaris", "wasip1", "windows", "zos"} {
				osSymbols = append(osSymbols, "GOOS_"+os)
			}
			return carapace.ActionValues(osSymbols...).Tag("target operating system symbols")
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			archSymbols := make([]string, 0)
			for _, arch := range []string{"386", "amd64", "amd64p32", "arm", "armbe", "arm64", "arm64be", "loong64", "mips", "mipsle", "mips64", "mips64le", "mips64p32", "mips64p32le", "ppc", "ppc64", "ppc64le", "riscv", "riscv64", "s390", "s390x", "sparc", "sparc64", "wasm"} {
				archSymbols = append(archSymbols, "GOARCH_"+arch)
			}
			return carapace.ActionValues(archSymbols...).Tag("target architecture symbols")
		}),
		carapace.ActionValuesDescribed(
			"GO386_sse2", "386 SSE2 extension",
			"GO386_softfloat", "386 soft float",
		).Tag("386 symbols"),
		carapace.ActionValuesDescribed(
			"GOAMD64_v1", "amd64 v1 microarchitecture",
			"GOAMD64_v2", "amd64 v2 microarchitecture",
			"GOAMD64_v3", "amd64 v3 microarchitecture",
			"GOAMD64_v4", "amd64 v4 microarchitecture",
		).Tag("amd64 symbols"),
		carapace.ActionValuesDescribed(
			"GOARM_5", "ARM v5",
			"GOARM_6", "ARM v6",
			"GOARM_7", "ARM v7",
		).Tag("arm symbols"),
		carapace.ActionValuesDescribed(
			"GOARM64_LSE", "ARM64 Large System Extensions",
		).Tag("arm64 symbols"),
		carapace.ActionValuesDescribed(
			"GOMIPS_hardfloat", "MIPS hard float",
			"GOMIPS_softfloat", "MIPS soft float",
		).Tag("mips symbols"),
		carapace.ActionValuesDescribed(
			"GOMIPS64_hardfloat", "MIPS64 hard float",
			"GOMIPS64_softfloat", "MIPS64 soft float",
		).Tag("mips64 symbols"),
		carapace.ActionValuesDescribed(
			"GOPPC64_power8", "PowerPC64 v8",
			"GOPPC64_power9", "PowerPC64 v9",
			"GOPPC64_power10", "PowerPC64 v10",
		).Tag("ppc64 symbols"),
		carapace.ActionValuesDescribed(
			"GORISCV64_rva22u64", "RISC-V 64 rva22u64",
			"GORISCV64_rv64gc", "RISC-V 64 rv64gc",
		).Tag("riscv64 symbols"),
		carapace.ActionValuesDescribed(
			"GOBUILDMODE_shared", "shared build mode",
		).Tag("build mode symbols"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			expSymbols := make([]string, 0)
			for _, exp := range []string{"arenas", "boringcrypto", "cgocheck2", "dwarf5", "fieldtrack", "goroutineleakprofile", "greenteagc", "heapminimum512kib", "jsonv2", "loopvar", "mapsplitgroup", "newinliner", "preemptibleloops", "randomizedheapbase64", "regabi", "regabiargs", "regabiwrappers", "runtimefreegc", "runtimesecret", "simd", "sizespecializedmalloc", "staticlockranking"} {
				expSymbols = append(expSymbols, "GOEXPERIMENT_"+exp)
			}
			return carapace.ActionValues(expSymbols...).Tag("experiment symbols")
		}),
	).ToA()
}

// ActionAsmPredefinedValues completes values for predefined assembler symbols
//
//	1 (default)
//	0
func ActionAsmPredefinedValues(symbol string) carapace.Action {
	return carapace.ActionValuesDescribed(
		"1", "enabled",
		"0", "disabled",
	).Tag("symbol values")
}
