package golang

import "github.com/rsteube/carapace"

// ActionOperatingSystems completes known operating systems
//
//	aix
//	android
func ActionOperatingSystems() carapace.Action {
	return carapace.ActionValues(
		"aix",
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"hurd",
		"illumos",
		"ios",
		"js",
		"linux",
		"nacl",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"wasip1",
		"windows",
		"zos",
	)
}

// ActionUnixOperatingSystems completes known operating systems matched by the "unix" build tag
//
//	darwin
//	linux
func ActionUnixOperatingSystems() carapace.Action {
	return carapace.ActionValues(
		"aix",
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"hurd",
		"illumos",
		"ios",
		"linux",
		"netbsd",
		"openbsd",
		"solaris",
	)
}

// ActionArchitectures completes known architectures
//
//	386
//	amd64
func ActionArchitectures() carapace.Action {
	return carapace.ActionValues(
		"386",
		"amd64",
		"amd64p32",
		"arm",
		"armbe",
		"arm64",
		"arm64be",
		"loong64",
		"mips",
		"mipsle",
		"mips64",
		"mips64le",
		"mips64p32",
		"mips64p32le",
		"ppc",
		"ppc64",
		"ppc64le",
		"riscv",
		"riscv64",
		"s390",
		"s390x",
		"sparc",
		"sparc64",
		"wasm",
	)
}
