package zig

import "github.com/carapace-sh/carapace"

// ActionFiles completes supported files.
//
//	main.zig
//	file.cxx
func ActionFiles() carapace.Action {
	return carapace.ActionFiles(
		".zig",                              // "Zig source code",
		".o",                                // "ELF object file",
		".o",                                // "Mach-O (macOS) object file",
		".o",                                // "WebAssembly object file",
		".obj",                              // "COFF (Windows) object file",
		".lib",                              // "COFF (Windows) static library",
		".a",                                // "ELF static library",
		".a",                                // "Mach-O (macOS) static library",
		".a",                                // "WebAssembly static library",
		".so",                               // "ELF shared object (dynamic link)",
		".dll",                              // "Windows Dynamic Link Library",
		".dylib",                            // "Mach-O (macOS) dynamic library",
		".tbd",                              // "(macOS) text-based dylib definition",
		".s",                                // "Target-specific assembly source code",
		".S",                                // "Assembly with C preprocessor (requires LLVM extensions)",
		".c",                                // "C source code (requires LLVM extensions)","
		".cxx", ".cc", ".C", ".cpp", ".c++", // "C++ source code (requires LLVM extensions)",
		".m",  // "Objective-C source code (requires LLVM extensions)",
		".mm", // "Objective-C++ source code (requires LLVM extensions)",
		".bc", // "LLVM IR Module (requires LLVM extensions)",
	)
}
